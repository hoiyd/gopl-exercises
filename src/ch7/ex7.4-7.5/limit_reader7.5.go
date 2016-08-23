package main

import (
	"fmt"
	"io"
	"strings"
)

type IOLimitReader struct {
	reader io.Reader
	n      int
}

func (lr *IOLimitReader) Read(p []byte) (int, error) {
	n, _ := lr.reader.Read(p)
	if n > lr.n {
		n = lr.n
	}
	return n, io.EOF
}

// Limit the # of bytes reader can read to n
func LimitReader(reader io.Reader, n int) *IOLimitReader {
	lr := IOLimitReader{reader, n}
	return &lr
}

func main() {
	r := LimitReader(strings.NewReader("<html><body><h1>hello</h1></body></html>aaaaaa"), 40)
	buffer := make([]byte, 1024)
	n, err := r.Read(buffer)
	buffer = buffer[:n]
	fmt.Println(n, err, string(buffer))
}
