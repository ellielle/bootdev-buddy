package main

import (
	"bytes"
	"io"

	"golang.org/x/net/html"
)

func traverseNodes(node *html.Node, matcher func(node *html.Node) (keep bool, exit bool)) (nodes []*html.Node) {
	var keep, exit bool
	var f func(*html.Node)

	f = func(node *html.Node) {
		keep, exit = matcher(node)
		if keep {
			nodes = append(nodes, node)
		}
		if exit {
			return
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			f(child)
		}
	}
	f(node)
	return nodes
}

func renderNode(node *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, node)
	return buf.String()
}
