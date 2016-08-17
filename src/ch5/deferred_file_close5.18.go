package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, url := range os.Args[1:] {
		if filename, n, err := fetchHTMLContentIntoFile(url); err == nil {
			fmt.Printf("%d bytes have been downloaded into %s\n", n, filename)
		} else {
			fmt.Println(err)
		}
	}
}

func fetchHTMLContentIntoFile(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	return local, n, err
}
