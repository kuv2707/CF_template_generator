package Parse

import (
	"cf_template_generator/langTemplates"
	"fmt"
	"html"
	"strings"
)

/*
*
 */
func ParseProblem(domstr string) {
	//set QUESTION_NAME and TESTCASETEMPLATE from domstr

	root := CreateDom(domstr)
	ps := findNodeInDOM(root, "class", "problem-statement")
	if ps == nil {
		fmt.Println("Error in finding input-specification in question document")
		return
	}
	title := findNodeInDOM(ps, "class", "title")
	langTemplates.TemplateInfo["QUESTION_NAME"] = slugify(html.UnescapeString(title.FirstChild.Data))


	fmt.Println("title:", langTemplates.TemplateInfo["QUESTION_NAME"])

	inp := findNodeInDOM(ps, "class", "input-specification")

	strcontent := allStrContentInDOM(inp)
	includeLoop := strings.Contains(strcontent, "test cases")
	if includeLoop {
		langTemplates.TemplateInfo["TESTCASETEMPLATE"] = langTemplates.Syntaxes[langTemplates.TemplateInfo["lang"]][0]
	} else {
		langTemplates.TemplateInfo["TESTCASETEMPLATE"] = langTemplates.Syntaxes[langTemplates.TemplateInfo["lang"]][1]
	}

	return
}


func slugify(s string)string{
	delimiters:="(). "
	ret:=""
	bucket:=""
	for i:=0;i<len(s);i++{
		if strings.ContainsAny(delimiters,string(s[i])){
			if bucket==""{
				continue
			}
			ret+=bucket+"_"
			bucket=""
		}else{
			bucket+=string(s[i])
		}
	}
	if bucket!=""{
		ret+=bucket
	}
	return ret
}
