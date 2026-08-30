package helpers

import (
	"fmt"
	"strings"
	"unicode"
)

func ValidateASCII(text string) bool {
	for pos, r := range text {
		if r > 127 {
			fmt.Printf(
				"Error: Invalid character '%c' (Unicode: U+%04X) at position %d. Only standard English ASCII text allowed.\n",
				r, r, pos,
			)
			return false
		}
	}
	return true
}

func capitalize(word string) string {
	if word == "" {
		return ""
	}

	runes := []rune(word)

	for i, r := range runes {
		if unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)

			for j := i + 1; j < len(runes); j++ {
				runes[j] = unicode.ToLower(runes[j])
			}
			return string(runes)
		}
	}
	return word
}

func isPunctuation(char rune) bool {
	return char == '.' || char == ',' ||
		char == '!' || char == '?' ||
		char == ':' || char == ';'
}

func NormalizeCommands(text string) string {
	for {
		newText := strings.ReplaceAll(text, "(cap ", "(cap")
		newText = strings.ReplaceAll(newText, "(up ", "(up")
		newText = strings.ReplaceAll(newText, "(low ", "(low")

		if newText == text {
			break
		}
		text = newText
	}
	return text
}

func NormalizeWords(words []string) []string {
	result := []string{}

	for i := 0; i < len(words); i++ {
		if i+1 < len(words) &&
			(words[i] == "(cap" || words[i] == "(up" || words[i] == "(low") &&
			words[i+1] == "," {

			result = append(result, words[i]+",")
			i++
			continue
		}
		result = append(result, words[i])
	}
	return result
}

func startsWithVowelOrH(word string) bool {
	word = strings.Trim(word, `"'()`)

	if word == "" {
		return false
	}

	first := strings.ToLower(word)[0]

	return first == 'a' || first == 'e' ||
		first == 'i' || first == 'o' ||
		first == 'u' || first == 'h'
}

func FixArticle(words []string) {
	for i := 0; i < len(words)-1; i++ {
		word := words[i]

		// Normal a / A
		if word == "a" || word == "A" {
			if startsWithVowelOrH(words[i+1]) {
				if word == "a" {
					words[i] = "an"
				} else {
					words[i] = "An"
				}
			}
		}

		// (a / (A
		if word == "(a" || word == "(A" {
			nextWord := strings.TrimRight(words[i+1], ")")

			if startsWithVowelOrH(nextWord) {
				if word == "(a" {
					words[i] = "(an"
				} else {
					words[i] = "(An"
				}
			}
		}

		//'a 'A
		if word == "'a" || word == "'A" {
			nextWord := strings.TrimRight(words[i+1], ")")

			if startsWithVowelOrH(nextWord) {
				if word == "'a" {
					words[i] = "'an"
				} else {
					words[i] = "'An"
				}
			}
		}
	}
}

func isApostrophe(word string, index int) bool {
	if word[index] != '\'' {
		return false
	}
	return index > 0 && index < len(word)-1
}

func hasLetter(word string) bool {
	for _, char := range word {
		if (char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') {
			return true
		}
	}
	return false
}
