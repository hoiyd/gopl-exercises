package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	p := make(chan string)
	q := make(chan string)
	var i int64

	start := time.Now()

	go func() {
		p <- "ping"
		for {
			i++
			<-q
			p <- "ping"
		}
	}()

	go func() {
		for {
			<-p
			q <- "pong"
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	timeSpent := time.Since(start).Seconds()
	fmt.Printf("count=%d, timeSpent = %v\n", i, timeSpent)
	fmt.Printf("%v roundtrips per second.\n", float64(i)/float64(timeSpent))
}
