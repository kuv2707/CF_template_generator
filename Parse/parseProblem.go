package Parse

import (
	"fmt"

	"golang.org/x/net/html"
)

/**


 */
func ParseProblem(domstr string) map[string]string {
	root:=CreateDom(domstr)
	fmt.Println("after dom creation",root.Type)
	printDOM(traverseDOM(root))
}