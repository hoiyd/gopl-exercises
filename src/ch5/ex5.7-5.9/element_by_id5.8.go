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
		node := ElementByID(doc, "footer")
		// for _, a := range node.Attr {
		// 	fmt.Printf("%s = %s\n", a.Key, a.Val)
		// }
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

func ElementByID(doc *html.Node, id string) *html.Node {
	var node *html.Node

	pre := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					node = n
					return true
				}
			}
		}
		return false
	}

	forEachDoc(doc, pre, nil)
	return node
}

func forEachDoc(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil {
		if found := pre(n); found {
			// If found then return immediately
			return true
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if found := forEachDoc(c, pre, post); found {
			return true
		}
	}

	if post != nil {
		if found := post(n); found {
			return true
		}
	}
	return false
}
