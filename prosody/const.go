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

var DEERGHA_SIGNS = []string{
	"ा",
	"ी",
	"ू",
	"ॄ",
	"ॅ",
	"े",
	"ै",
	"ॉ",
	"ो",
	"ौ",
	"ं",
	"ः",
	"ँ",
}

var HRASVA_SIGNS = []string{
	"ि",
	"ु",
	"ृ",
}

var ENDING_MARKS = []string{
	"ं",
	"ः",
	"ँ",
}

var DEERGHA_VOWELS = []string{
	"आ",
	"ई",
	"ऊ",
	"ॠ",
	"ए",
	"ऐ",
	"ओ",
	"औ",
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
