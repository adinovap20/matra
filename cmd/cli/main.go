package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/adinovap20/matra/prosody"
)

func main() {
	inputFile := flag.String("i", "input.txt", "Input file")
	outputFile := flag.String("o", "output.json", "Output file")
	flag.Parse()

	fmt.Println("Input file:", *inputFile)
	fmt.Println("Output file:", *outputFile)

	bytes, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("Failed to read from file: %s", err)
	}
	content := string(bytes)

	result := prosody.Analyze(content)
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal json: %v", err)
	}
	err = os.WriteFile(*outputFile, data, 0644)

}
