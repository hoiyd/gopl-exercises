package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	prefix = "http://"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("HTTP response status code: %s\n", resp.Status)

		// The first return value of io.Copy is the nuber of bytes that are copied,
		// which in this case, is useless
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
