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
	writer := make(chan string) // outgoing client messages
	go clientWriter(conn, writer)

	reader := make(chan string)
	go clientReader(conn, reader)

	var who string
	nameTimer := time.NewTimer(timeout)
	writer <- "Enter your name:"
	select {
	case nickname := <-reader:
		who = nickname
	case <-nameTimer.C:
		conn.Close()
	}

	cli := client{name: who, channel: writer}

	writer <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	timer := time.NewTimer(timeout)

	// 如果不加for循环，只要发送一条消息就会被select到，然后导致select解除阻塞，运行后续的退出操作。
	// 这不是我们预期的，所以需要加上for循环与select配合使用。
Loop:
	for {
		select {
		case msg := <-reader:
			messages <- who + ": " + msg
			timer.Reset(timeout)
		case <-timer.C:
			conn.Close()
			break Loop
		}
	}

	leaving <- cli
	messages <- who + " has left"
	timer.Stop()
	conn.Close()
}

func clientWriter(conn net.Conn, writer <-chan string) {
	for msg := range writer {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func clientReader(conn net.Conn, reader chan<- string) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		reader <- input.Text()
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
