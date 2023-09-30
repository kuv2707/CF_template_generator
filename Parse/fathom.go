package Parse

import (
	lts "cf_template_generator/langTemplates"
	"fmt"
	"os"
	"strings"
)

var url ="https://codeforces.com/problemset/problem/%s/%s"
//1578/C

/**
 * @brief      { function_description }
 *
 * @param      args  The arguments
 *
 * @return     {  }
 */
func FathomArgsAndValidate(args []string)(string, string) {

	/*

	*/
	lang:=""
	if len(args)<2{
		fmt.Println("No arguments provided")
		os.Exit(-1)
	}
	if(len(args)<3){
		fmt.Println("Language not provided. Defaulting to cpp")
		lang="cpp"
	}
	qstn:=args[1]
	namenum:=strings.Split(qstn, "/")
	if(len(namenum)!=2){
		fmt.Println("Invalid question name")
		os.Exit(-1)
	}

	validurl:=fmt.Sprintf(url,namenum[0],namenum[1])
	if !strings.EqualFold(lang,"cpp"){
		
		if !lts.IsAllowedLanguage(args[2]){
			fmt.Println("This language is not supported\nSupported languages are:",lts.AllowedLanguages,"\nDefaulting to cpp")
			lang="cpp"
		}else{
			lang=args[2]
		}
	}




	return validurl,lang
}

