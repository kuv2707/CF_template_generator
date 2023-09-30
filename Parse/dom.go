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

func findNodeInDOM(node *html.Node,key string, value string) *html.Node {
	if node.Type == html.ElementNode && node.FirstChild != nil {
		for _, attr := range node.Attr {
			// fmt.Printf("Attribute: %s=\"%s\"\n", attr.Key, attr.Val)
			if attr.Key == key && attr.Val == value {
				return node
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		n := findNodeInDOM(child,key,value)
		if n != nil {
			return n
		}

	}
	return nil
}


func allStrContentInDOM(node *html.Node)string{
	if node.Type == html.TextNode {
		return node.Data
	}
	ret:=""
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		ret+=allStrContentInDOM(child)
	}
	return ret
}