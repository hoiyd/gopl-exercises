package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		doc, err := fetchHTMLContent(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch html content: %v\n", err)
			continue
		}
		// nodes := ElementByNames(doc, "img")
		nodes := ElementByNames(doc, "h1", "h2", "h3", "h4", "input")
		fmt.Printf("In total %d elements are found.\n", len(nodes))
		for i, node := range nodes {
			fmt.Printf("Element #%d:\n", i+1)
			fmt.Printf("<%s", node.Data)
			for _, a := range node.Attr {
				fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
			}
			fmt.Printf(">\n")
		}
	}
}

func fetchHTMLContent(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return doc, err
}

func ElementByNames(doc *html.Node, names ...string) []*html.Node {
	var nodes []*html.Node

	pre := func(n *html.Node) {
		if n.Type == html.ElementNode && includes(names, n.Data) {
			nodes = append(nodes, n)
		}
	}

	forEachDoc(doc, pre, nil)
	return nodes
}

func forEachDoc(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachDoc(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func includes(strings []string, s string) bool {
	for _, str := range strings {
		if str == s {
			return true
		}
	}
	return false
}
