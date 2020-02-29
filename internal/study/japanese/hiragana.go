package japanese

var hiraganaKeys []string

var hiragana = map[string]string{
	"あ": "a",
}

func init() {
	for k := range hiragana {
		hiraganaKeys = append(hiraganaKeys, k)
	}
}
