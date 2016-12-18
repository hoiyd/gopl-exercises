package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	resetTimerCh := make(chan string)
	go countTime(c, resetTimerCh)

	for input.Scan() {
		resetTimerCh <- "Reset timer"
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func countTime(conn net.Conn, resetTimerCh <-chan string) {
	timer := time.NewTicker(1 * time.Second)
	counter := 0
	max := 10
	for {
		select {
		case <-timer.C:
			counter++
			if counter >= max {
				msg := conn.RemoteAddr().String() + " no input for too long. Client is now kicked out."
				fmt.Println(msg)
				fmt.Fprintln(conn, msg) // Let to-be-closed client see this msg
				timer.Stop()
				conn.Close()
				return
			}
		case <-resetTimerCh:
			counter = 0 //reset counter/timer
		}
	}
}
