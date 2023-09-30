package Parse

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)
func CreateDom(domstr string)*html.Node{
	doc,err:=html.Parse(strings.NewReader(domstr))
	if(err != nil){
		panic(err)
	}
	return doc

}

func printDOM(node *html.Node) {
	if node == nil {
		return
	}
	if node.Type == html.ElementNode && node.FirstChild != nil {
		fmt.Printf("Element: <%s>\n", node.FirstChild.Data)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		printDOM(child)
	}
}

func traverseDOM(node *html.Node) *html.Node {
	if node.Type == html.ElementNode && node.FirstChild != nil {
		for _, attr := range node.Attr {
			// fmt.Printf("Attribute: %s=\"%s\"\n", attr.Key, attr.Val)
			if attr.Key == "class" && attr.Val == "problem-statement" {
				fmt.Println("found")
				return node
			}
		}

	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		n := traverseDOM(child)
		if n != nil {
			return n
		}

	}
	return nil
}
