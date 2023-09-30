package langTemplates

var AllowedLanguages =[5]string{"cpp","java","python","go"}
var DefaultTemplate = map[string]string{
}

func IsAllowedLanguage(lang string)bool{
	for _,l:=range AllowedLanguages{
		if l==lang{
			return true
		}
	}
	return false
}