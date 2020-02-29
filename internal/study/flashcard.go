package study

// Flashcard is a flashcard man, come on
type Flashcard struct {
	Name   string
	Prompt string
	Answer string
}

// Eval checks wether the given answer is correct for the flashcard
func (f Flashcard) Eval(ans string) bool {
	return f.Prompt == ans
}
