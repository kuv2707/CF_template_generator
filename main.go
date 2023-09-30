package main

import (
	"cf_template_generator/Parse"
	"cf_template_generator/langTemplates"
	"fmt"
	"os"
)

func main() {
	url, lang := Parse.FathomArgsAndValidate(os.Args)
	fmt.Println("problem url: ", url)
	fmt.Println("chosen lang: ", lang)
	

	langTemplates.TemplateInfo["QUESTION_URL"] = url
	langTemplates.TemplateInfo["lang"] = lang
	langTemplates.TemplateInfo["TESTCASETEMPLATE"] = "while(false){}"

	domstr, err := Parse.ApiCall(url)
	if err != nil {
		fmt.Println("Error in getting question info: ", err,"\n Creating default template")
		langTemplates.CreateTemplate()
		return
	}
	// langTemplates.WriteToFile("dom.html", domstr)

	Parse.ParseProblem(domstr)
	langTemplates.CreateTemplate()

}
