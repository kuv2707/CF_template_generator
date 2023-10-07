package langTemplates

var AllowedLanguages = [5]string{"cpp", "java", "python", "go"}

func IsAllowedLanguage(lang string) bool {
	for _, l := range AllowedLanguages {
		if l == lang {
			return true
		}
	}
	return false
}

var Syntaxes = map[string][]string{
	"cpp": {`ll t;
    cin >> t;
    while(t--)
        solve();`,
		"solve();"},
	"java":   {"int tsc=sc.nextInt();\n			while(tsc-->0){}", ""},
	"python": {"for _ in range(int(input())):", ""},
	"go":     {"for tsc:=scanInt();tsc>0;tsc--{}", ""},
}
