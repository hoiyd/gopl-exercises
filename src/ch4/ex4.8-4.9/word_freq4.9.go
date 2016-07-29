package main

import (
	"bufio"
	"fmt"
	// "io"
	"os"
	// "unicode"
	// "unicode/utf8"
)

func main() {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()
		counts[word]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "word freq: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("word\tfreq\n")
	for word, n := range counts {
		fmt.Printf("%q\t%d\n", word, n)
	}
}
