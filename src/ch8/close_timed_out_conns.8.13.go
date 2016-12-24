package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

var timeout = 10 * time.Second

//!+broadcaster
type client struct {
	channel chan<- string // an outgoing message channel
	name    string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.channel <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			cli.channel <- "Here are the online clients at the moment:"
			for existingClient := range clients {
				cli.channel <- existingClient.name
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.channel)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{name: who, channel: ch}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	timer := time.NewTimer(timeout)

	// close掉conn之后，会导致input.Scan调用解除阻塞，从而继续运行leaving等清理工作。
	go func() {
		<-timer.C
		conn.Close()
	}()

	//在不加timeout机制之前，input.Scan()循环会永远阻塞在这里，
	//直到netcat客户端被kill掉，conn被关闭为止。
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		timer.Reset(timeout)
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- who + " has left"
	timer.Stop()
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
