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
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		// Look for link in <a> tag attributes
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val)
				links = append(links, a.Val)
			}
		}
	}
	links = visit(links, n.NextSibling) //first recursively visit its siblings
	links = visit(links, n.FirstChild)  //then recursively visit its first child

	return links
}
