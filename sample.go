package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func resfes() {
	domStr := `
	<div>Div content is thisssss</div>
	`

	root := ParseCreateDOM(domStr)
	fmt.Println("After DOM creation:", root.Type)
	traversedom(root)
}

func ParseCreateDOM(domStr string) *html.Node {
	reader := strings.NewReader(domStr)
	doc, _ := html.Parse(reader)
	return doc
}

func traversedom(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "div" {
		fmt.Println("Div content:", node.FirstChild.Data)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		traversedom(child)
	}
}
