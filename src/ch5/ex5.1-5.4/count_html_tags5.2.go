package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stdout, "findlink1: %v\n", err)
		os.Exit(1)
	}

	count := make(map[string]int)
	count = visit(count, doc)
	// count = visitLoop(count, doc)
	fmt.Println("element\t\tcount")
	for key, value := range count {
		fmt.Printf("%q\t\t%d\n", key, value)
	}
}

func visit(count map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return count
	}
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	count = visit(count, n.NextSibling) //first recursively visit its siblings
	count = visit(count, n.FirstChild)  //then recursively visit its first child

	return count
}

func visitLoop(count map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count = visit(count, c)
	}
	return count
}
