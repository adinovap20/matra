//go:build js && wasm

package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/adinovap20/matra/prosody"
)

func analyze(this js.Value, args []js.Value) any {
	if len(args) == 0 {
		return `{"error":"Missing input text"}`
	}
	result := prosody.Analyze(args[0].String())

	data, err := json.Marshal(result)
	if err != nil {
		return `{"error":"Could not encode result"}`
	}

	return string(data)
}

func main() {
	js.Global().Set(
		"analyzeSanskrit",
		js.FuncOf(analyze),
	)

	select {}
}
