// Use a WaitGroup to determine when the work is done, the `tokens` chan as a
// semaphore to limit concurrent requests, and a mutex around the `seen` map to
// avoid concurrent reads and writes.
package main

import (
	"./links"
	"flag"
	"fmt"
	"log"
	"sync"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)
var maxDepth int
// In the examples, this map is accessed by main routine only.
// But here it is accessed by many different go routines, so it requires
// mutex to pretect it.
var seen = make(map[string]bool)
var seenLock = sync.Mutex{}

func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(depth, url)
	if depth >= maxDepth {
		return
	}
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	for _, link := range list {
		seenLock.Lock()
		if seen[link] {
			seenLock.Unlock()
			continue
		}
		seen[link] = true
		seenLock.Unlock()
		wg.Add(1)
		go crawl(link, depth+1, wg)
	}
}

func main() {
	flag.IntVar(&maxDepth, "d", 3, "max crawl depth")
	flag.Parse()
  fmt.Printf("maxDepth: %d\n", maxDepth)

  wg := &sync.WaitGroup{}

  link :=flag.Args[0]
		wg.Add(1)
		go crawl(link, 0, wg)
	}

  wg.Wait()
	fmt.Println("Crawl finished.")
}
