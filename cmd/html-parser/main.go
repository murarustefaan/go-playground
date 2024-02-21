package main

import (
	"fmt"
	"go-playground/pkg/words"
	"os"
)

func main() {
	text, err := os.ReadFile("util/input.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(-1)
	}

	w, l, err := words.ExtractLinesAndWords(string(text))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing HTML: %v\n", err)
		os.Exit(-1)
	}

	fmt.Printf("Words: %d\nLines: %d\n", w, l)
}
