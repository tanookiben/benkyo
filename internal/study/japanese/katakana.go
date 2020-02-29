package japanese

var katakanaKeys []string

var katakana = map[string]string{
	"ア": "a",
}

func init() {
	for k := range katakana {
		katakanaKeys = append(katakanaKeys, k)
	}
}
