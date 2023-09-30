package main

import (
	Parse "cf_template_generator/Parse"
	"cf_template_generator/langTemplates"
	"fmt"
	"os"
)

func main() {
	url, lang := Parse.FathomArgsAndValidate(os.Args)
	fmt.Println("url: ", url)
	fmt.Println("lang: ", lang)
	if true {
		return
	}
	domstr, err := Parse.ApiCall(url)
	if err != nil {
		fmt.Println("Error: ", err," Creating default template for ",lang)
		langTemplates.CreateTemplate(langTemplates.DefaultTemplate)
		return
	}

	// domstr:=`
	// <div class="problem-statement"><h1>This is the heading</h1></div>
	// `

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)// if cant access pwd, then cant write to file either
		return
	}
	
	problem_info:=Parse.ParseProblem(domstr)
	
	
	fmt.Println("Creating template for ",lang," at ",pwd)

}


