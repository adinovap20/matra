package prosody

import "slices"

const (
	HALANT = "्"
)

var VOWEL_SIGNS = []string{
	"ा",
	"ि",
	"ी",
	"ु",
	"ू",
	"ृ",
	"ॄ",
	"ॅ",
	"ॆ",
	"े",
	"ै",
	"ॉ",
	"ॊ",
	"ो",
	"ौ",
}

var ENDING_MARKS = []string{
	"ं",
	"ः",
	"ँ",
}

func isVowelSignOrVirama(ch string) bool {
	if slices.Contains(VOWEL_SIGNS, ch) {
		return true
	}
	if ch == HALANT {
		return true
	}
	return false
}

func isEndingMark(ch string) bool {
	return slices.Contains(ENDING_MARKS, ch)
}
