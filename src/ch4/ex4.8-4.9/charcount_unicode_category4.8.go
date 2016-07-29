package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[string]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
		case unicode.IsControl(r):
			counts["control"]++
		case unicode.IsLetter(r):
			counts["letter"]++
		case unicode.IsMark(r):
			counts["mark"]++
		case unicode.IsNumber(r):
			counts["number"]++
		case unicode.IsPunct(r):
			counts["punct"]++
		case unicode.IsSpace(r):
			counts["space"]++
		case unicode.IsSymbol(r):
			counts["symbol"]++
		}
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)

	}

	fmt.Printf("\n")

	fmt.Printf("nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Println("\n%d invalid UTF-8 characters\n", invalid)
	}
}
