package prosody

import (
	"strings"
)

func analyzeLine(line string, idx int) LineAnalysis {
	normalizedLine := normalizeLine(line)
	aksharas := getAksharas(normalizedLine)
	matras := getMatras(aksharas)
	matraCnt := 0
	ganas := getGanas(matras)
	for _, c := range matras {
		matraCnt += c
	}
	marathiVrutta, _ := IdentifyMarathiVrutta(matras)
	sanskritVrutta, _ := IdentifySanskritVrutta(matras)
	lineAnalysis := LineAnalysis{
		Index:          idx,
		Line:           line,
		NormalizedLine: normalizedLine,
		Aksharas:       aksharas,
		AksharasCount:  len(aksharas),
		Matras:         matras,
		MatraCount:     matraCnt,
		Ganas:          ganas,
		MarathiVrutta:  marathiVrutta,
		SanskritVrutta: sanskritVrutta,
	}
	return lineAnalysis
}

func getAksharas(normalizedLine string) []string {
	aksharas := []string{}
	runes := []rune(normalizedLine)
	for i := 0; i < len(runes); {
		var sb strings.Builder
		sb.WriteRune(runes[i])
		i++

		if i < len(runes) && string(runes[i]) == HALANT {
			sb.WriteRune(runes[i])
			i++

			if i < len(runes) {
				sb.WriteRune(runes[i])
				i++
			}
		}

		// Attach vowel signs, halant, visarga, anusvara, and chandrabindu to the current akshara
		for i < len(runes) && (isVowelSignOrVirama(string(runes[i])) || isEndingMark(string(runes[i]))) {
			sb.WriteRune(runes[i])
			i++
		}

		// if the next consonant is followed by halant, attach both to the current akshara
		if i+1 < len(runes) && string(runes[i+1]) == HALANT {
			sb.WriteRune(runes[i])
			sb.WriteRune(runes[i+1])
			i += 2
		}

		aksharas = append(aksharas, sb.String())
	}
	return aksharas
}

func getMatras(aksharas []string) []int {
	matras := []int{}
	for _, akshara := range aksharas {
		if strings.HasSuffix(akshara, HALANT) {
			matras = append(matras, 2)
			continue
		}

		foundMatra := false
		for _, deerghaSign := range DEERGHA_SIGNS {
			if strings.HasSuffix(akshara, deerghaSign) {
				foundMatra = true
			}
		}
		if foundMatra {
			matras = append(matras, 2)
			continue
		}

		foundMatra = false
		for _, deerghaVowel := range DEERGHA_VOWELS {
			if akshara == deerghaVowel {
				foundMatra = true
			}
		}
		if foundMatra {
			matras = append(matras, 2)
			continue
		}

		matras = append(matras, 1)
	}
	return matras
}

func getGanas(matras []int) string {
	var sb strings.Builder
	for ptr := 0; ptr < len(matras); {
		if ptr+2 < len(matras) {
			if matras[ptr] == 1 && matras[ptr+1] == 2 && matras[ptr+2] == 2 {
				sb.WriteRune('य')
			} else if matras[ptr] == 2 && matras[ptr+1] == 2 && matras[ptr+2] == 2 {
				sb.WriteRune('म')
			} else if matras[ptr] == 2 && matras[ptr+1] == 2 && matras[ptr+2] == 1 {
				sb.WriteRune('त')
			} else if matras[ptr] == 2 && matras[ptr+1] == 1 && matras[ptr+2] == 2 {
				sb.WriteRune('र')
			} else if matras[ptr] == 1 && matras[ptr+1] == 2 && matras[ptr+2] == 1 {
				sb.WriteRune('ज')
			} else if matras[ptr] == 2 && matras[ptr+1] == 1 && matras[ptr+2] == 1 {
				sb.WriteRune('भ')
			} else if matras[ptr] == 1 && matras[ptr+1] == 1 && matras[ptr+2] == 1 {
				sb.WriteRune('न')
			} else if matras[ptr] == 1 && matras[ptr+1] == 1 && matras[ptr+2] == 2 {
				sb.WriteRune('स')
			}
			ptr += 3
			continue
		}
		if matras[ptr] == 1 {
			sb.WriteRune('ल')
			ptr++
			continue
		}
		sb.WriteString("गा")
		ptr++
	}
	return sb.String()
}

func normalizeLine(line string) string {
	stringsToReplace := []string{
		" ", "\t", "\r", "\n",
		"॥", "।",
		"१", "२", "३", "४", "५", "६", "७", "८", "९", "०",
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
		"[", ".", ",", ";", ":", "!", "?", "-", "—", "'", "\"", "“", "”", "‘", "’",
		"(", ")", "[", "]", "{", "}", "]",
	}
	normalized := line
	for _, sr := range stringsToReplace {
		normalized = strings.ReplaceAll(normalized, sr, "")
	}
	return normalized
}
