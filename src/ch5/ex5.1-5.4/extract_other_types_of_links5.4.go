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
	visit(doc)
}

func visit(n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode && linkExtractable(n) {
		for _, a := range n.Attr {
			if a.Key == "src" || a.Key == "href" {
				fmt.Println(a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}

// Other elements with links other than <a>
func linkExtractable(n *html.Node) bool {
	return n.Data == "a" ||
		n.Data == "img" ||
		n.Data == "script" ||
		n.Data == "style" ||
		n.Data == "link"
}
