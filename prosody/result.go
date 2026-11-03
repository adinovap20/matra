package prosody

import (
	"regexp"
)

type LineAnalysis struct {
	Index          int      `json:"index"`
	Line           string   `json:"line"`
	NormalizedLine string   `json:"normalizedLine"`
	Aksharas       []string `json:"aksharas"`
	AksharasCount  int      `json:"aksharasCount"`
	MatraCount     int      `json:"matraCount"`
	Matras         []int    `json:"matras"`
	Ganas          string   `json:"ganas"`
	MarathiVrutta  string   `json:"marathi_vrutta"`
	SanskritVrutta string   `json:"sanskrit_vrutta"`
}

type Result struct {
	Input    string         `json:"input"`
	Analysis []LineAnalysis `json:"analysis"`
}

func Analyze(text string) Result {
	re := regexp.MustCompile(`[\n॥।|]`)
	lines := re.Split(text, -1)
	lineAnalysis := []LineAnalysis{}
	for idx, line := range lines {
		if len(line) == 0 {
			continue
		}
		lineAnalysis = append(lineAnalysis, analyzeLine(line, idx+1))
	}
	return Result{
		Input:    text,
		Analysis: lineAnalysis,
	}
}
