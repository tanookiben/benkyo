package study

import (
	"math/rand"
	"time"
)

// Study is essentially a group of flashcards
type Study struct {
	Cards []Flashcard
	Order []int
	Index int
}

// NewStudy returns a new study with the specified flashcards
func NewStudy(cards []Flashcard) *Study {
	s := &Study{
		Cards: cards,
	}
	s.init()
	return s
}

// Eval checks if the answer is correct for the current card and increments the
// index true
func (s *Study) Eval(ans string) bool {
	if s.current().Eval(ans) {
		s.next()
		return true
	}
	return false
}

// Cheat returns the answer for the current card and increments the index
func (s *Study) Cheat() string {
	defer s.next()
	return s.current().Answer
}

func (s *Study) current() Flashcard {
	return s.Cards[s.Index]
}

func (s *Study) next() {
	s.Index++
	if s.Index == len(s.Cards) {
		s.Index = 0
		s.init()
	}
}

func (s *Study) init() {
	n := len(s.Cards)
	var order []int
	for i := 0; i < n; i++ {
		order = append(order, i)
	}
	shuffle(order)
	s.Order = order
}

func shuffle(a []int) {
	rand := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < len(a); i++ {
		t := a[i]
		n := rand.Intn(len(a))
		a[i] = a[n]
		a[n] = t
	}
}
