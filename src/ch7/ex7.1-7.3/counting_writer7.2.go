package main

import (
	"bytes"
	"fmt"
	"io"
	// "os"
)

type ByteWriter struct {
	w     io.Writer
	count int64
}

func (bw *ByteWriter) Write(p []byte) (int, error) {
	n, err := bw.w.Write(p)
	bw.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	writer := ByteWriter{w, 0}
	return &writer, &writer.count
}

func main() {
	var buf bytes.Buffer
	countingWriter, count := CountingWriter(&buf)

	// if I use Stdin, Stdout or Stderr, just like below
	// countingWriter, count := CountingWriter(os.Stderr)
	// "Hello world" would be printed out on the console, interesting.

	fmt.Fprint(countingWriter, "Hello world\n") // "Hello world" won't be printed out on the console.
	fmt.Println(*count)
}
