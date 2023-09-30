package langTemplates

var AllowedLanguages =[5]string{"cpp","java","python","go"}


func IsAllowedLanguage(lang string)bool{
	for _,l:=range AllowedLanguages{
		if l==lang{
			return true
		}
	}
	return false
}

func LoopSyntax(lang string)string{
	switch lang{
	case "cpp":
		return `ll t;
    cin >> t;
    while(t--)
        solve();`
	case "java":
		return "int tsc=sc.nextInt();\n			while(tsc-->0){}"
	case "python":
		return "for _ in range(int(input())):"
	case "go":
		return "for tsc:=scanInt();tsc>0;tsc--{}"
	default:
			return ""
	}
}
