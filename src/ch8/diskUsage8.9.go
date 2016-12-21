// Almost the same as du3.go

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var sema = make(chan struct{}, 20)
var verbose = flag.Bool("v", false, "show verbose progress messages.")

type SizeInfo struct {
	idx  int
	size int64
}

func walkDir(dir string, sizeInfo chan<- SizeInfo, idx int, n *sync.WaitGroup) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, sizeInfo, idx, n)
		} else {
			sizeInfo <- SizeInfo{idx, entry.Size()}
		}
	}
}

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dir: %s, du: %v\n", dir, err)
		return nil
	}
	return entries
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	sizeInfo := make(chan SizeInfo)
	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		go walkDir(root, sizeInfo, i, &n)
	}

	go func() {
		n.Wait()
		close(sizeInfo)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))
loop:
	for {
		select {
		case si, ok := <-sizeInfo:
			if !ok {
				break loop // 直接调到loop标签，但是不运行循环；如果goto则还会运行循环。
			}
			nfiles[si.idx]++
			nbytes[si.idx] += si.size
		case <-tick:
			printDiskUsage(roots, nfiles, nbytes)
		}
	}
	printDiskUsage(roots, nfiles, nbytes)
}

func printDiskUsage(roots []string, nfiles, nbytes []int64) {
	for i, r := range roots {
		fmt.Printf("%10d files  %.3f GB under %s\n", nfiles[i], float64(nbytes[i])/1e9, r)
	}
}
