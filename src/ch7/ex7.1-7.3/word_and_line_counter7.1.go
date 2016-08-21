package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int
type LineCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
	var count int
	input := bufio.NewScanner(strings.NewReader(string(p)))
	input.Split(bufio.ScanWords)

	for input.Scan() {
		count++
	}
	*wc = WordCounter(count)
	return count, nil
}

func (lc *LineCounter) Write(p []byte) (int, error) {
	var count int
	input := bufio.NewScanner(strings.NewReader(string(p)))
	input.Split(bufio.ScanLines)

	for input.Scan() {
		count++
	}
	*lc = LineCounter(count)
	return count, nil
}

func main() {
	var wc WordCounter
	var lc LineCounter

	name := "Will Smith"
	fmt.Fprintf(&wc, "hello, %s", name)
	fmt.Println(wc)

	fmt.Fprintf(&lc, "hello\n%s\n", name)
	fmt.Println(lc)
}
