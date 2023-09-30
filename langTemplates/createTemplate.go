package langTemplates

import (
	"bytes"
	"embed"
	"fmt"
	_ "io"
	"os"
)

var TemplateInfo = map[string]string{
	"directory": "/",
	"lang":"cpp",//redundant here?
	"QUESTION_NAME":"codeforces_qstn",
}


/*
*
creates a template at cwd
*/

//go:embed templates/*
var templates embed.FS

func CreateTemplate() {
	templatecontent,err:=templates.ReadFile("templates/"+TemplateInfo["lang"]+".template")
	if err!=nil{
		fmt.Println("Fatal error in reading template file: ",err,"\nFile is not created.")
		os.Exit(-1)
	}
	for key,value:=range TemplateInfo{
		templatecontent=bytes.ReplaceAll(templatecontent,[]byte(key),[]byte(value))
	}
	err=WriteToFile(TemplateInfo["directory"]+TemplateInfo["QUESTION_NAME"]+"."+TemplateInfo["lang"],string(templatecontent))
	if err!=nil{
		fmt.Println("Fatal error in writing template file: ",err,"\n File is not created.\n Exiting...")
		os.Exit(-1)
	}


	fmt.Println("Template created successfully!")
	return

}

func WriteToFile(filepath string, data string) error {
	
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err) // if cant access pwd, then cant write to file either
		return nil
	}

	f, err := os.Create(cwd+filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}
