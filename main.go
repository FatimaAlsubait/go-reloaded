package main

import (
	"fmt"
	helpers "go-reloaded/flags"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if !strings.HasSuffix(inputFile, ".txt") || !strings.HasSuffix(outputFile, ".txt") {
		fmt.Println("Error: files must be .txt files")
		return
	}
	// Read the input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	if len(strings.TrimSpace(string(data))) == 0 {
		fmt.Println("Error: input file is empty")
		return
	}

	text := string(data)
	text = strings.ReplaceAll(text, "(cap)", "(cap) ")
	text = strings.ReplaceAll(text, "(up)", "(up) ")
	text = strings.ReplaceAll(text, "(low)", "(low) ")
	text = strings.ReplaceAll(text, "(hex)", "(hex) ")
	text = strings.ReplaceAll(text, "(bin)", "(bin) ")

	text = strings.ReplaceAll(text, "(cap))", "(cap) )")
	text = strings.ReplaceAll(text, "(up))", "(up) )")
	text = strings.ReplaceAll(text, "(low))", "(low) )")
	text = strings.ReplaceAll(text, "(hex))", "(hex) )")
	text = strings.ReplaceAll(text, "(bin))", "(bin) )")

	text = strings.ReplaceAll(text, "(up ,", "(up, ")
	text = strings.ReplaceAll(text, "(low ,", "(low, ")
	text = strings.ReplaceAll(text, "(cap ,", "(cap, ")

	words := strings.Fields(text)
	words = helpers.NormalizeWords(words)
	
	// Apply the text modifications
	result, err := helpers.ProcessWords(words)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Fix "a" to "an"
	helpers.FixArticle(result)
	// Format punctuation and quotes
	finalText := helpers.FormatPunctuation(result)

	// Write the result to the output file
	err = os.WriteFile(outputFile, []byte(finalText), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
