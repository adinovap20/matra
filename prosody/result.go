package prosody

type Result struct {
	Input    string   `json:"input"`
	Aksharas []string `json:"aksharas"`
}

func Analyze(text string) Result {
	return Result{
		Input:    text,
		Aksharas: []string{"Hello", "world"},
	}
}
