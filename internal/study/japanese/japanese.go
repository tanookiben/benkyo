package japanese

import (
	"math/rand"
	"time"
)

const (
	// HiraganaStudy is the key for studying Hiragana
	HiraganaStudy = "hiragana"
	// KatakanaStudy is the key for studying Katakana
	KatakanaStudy = "katakana"
)

var (
	r *rand.Rand
)

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

// Prompt returns a random prompt for the specified charset
func Prompt(charset string) string {
	switch charset {
	case HiraganaStudy:
		return hiraganaKeys[rand.Intn(len(hiraganaKeys))]
	case KatakanaStudy:
		return katakanaKeys[rand.Intn(len(katakanaKeys))]
	default:
		return ""
	}
}

// Check compares the given answer with the expected answer for the prompt for
// the specified charset
func Check(charset, prompt, answer string) bool {
	switch charset {
	case HiraganaStudy:
		return hiragana[prompt] == answer
	case KatakanaStudy:
		return katakana[prompt] == answer
	default:
		return false
	}
}

// Hint returns the answer for the prompt for the specified charset
func Hint(charset, prompt string) string {
	switch charset {
	case HiraganaStudy:
		return hiragana[prompt]
	case KatakanaStudy:
		return katakana[prompt]
	default:
		return ""
	}
}
