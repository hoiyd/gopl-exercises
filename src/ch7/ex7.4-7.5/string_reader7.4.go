package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

type StringReader struct {
	str string
}

func (sr *StringReader) Read(p []byte) (int, error) {
	copy(p, []byte(sr.str))
	return len(sr.str), io.EOF
}

func NewReader(s string) *StringReader {
	sr := StringReader{s}
	return &sr
}

func main() {
	doc, _ := html.Parse(NewReader("<html><body><h1>hello</h1></body></html>"))
	fmt.Println(doc.FirstChild.LastChild.FirstChild.FirstChild.Data)
}
