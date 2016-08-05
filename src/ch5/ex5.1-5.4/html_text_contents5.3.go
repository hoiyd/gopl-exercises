package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stdout, "printHTMLText: %v\n", err)
		os.Exit(1)
	}

	printHTMLText(doc)
}

func printHTMLText(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data != "style" && c.Data != "script" && c.Data != "noscript" {
			printHTMLText(c)
		}
	}
}
