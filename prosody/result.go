package prosody

import "strings"

type LineAnalysis struct {
	Index          int      `json:"index"`
	Line           string   `json:"line"`
	NormalizedLine string   `json:"normalizedLine"`
	Aksharas       []string `json:"aksharas"`
	AksharasCount  int      `json:"aksharasCount"`
}

type Result struct {
	Input    string         `json:"input"`
	Analysis []LineAnalysis `json:"analysis"`
}

func Analyze(text string) Result {
	lines := strings.Split(text, "\n")
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
