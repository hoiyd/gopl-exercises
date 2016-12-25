package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	cancel := make(chan struct{})
	responses := make(chan *http.Response)
	wg := &sync.WaitGroup{}
	for _, url := range os.Args[1:] {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			req, err := http.NewRequest("HEAD", url, nil)
			if err != nil {
				log.Printf("HEAD %s: %s", url, err)
				return
			}

			req.Cancel = cancel
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("HEAD %s: %s", url, err)
				return
			}
			responses <- resp
		}(url)
	}
	resp := <-responses
	defer resp.Body.Close()
	close(cancel)
	fmt.Println(resp.Request.URL)
	for name, vals := range resp.Header {
		fmt.Printf("%s: %s\n", name, strings.Join(vals, ","))
	}
	// 不wai的话，main goroutine会立刻运行完毕而退出，从而导致整个程序退出。
	// 所以必须要在这里wait以阻塞main goroutine，使其有足够的时间等待其他goroutine的运行结果。
	wg.Wait()
}
