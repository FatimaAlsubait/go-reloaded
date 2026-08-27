package helpers

import "strings"

func capitalize(word string) string {
	if word == "" {
		return ""
	}

	return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
}

func isPunctuation(char rune) bool {
	return char == '.' || char == ',' ||
		char == '!' || char == '?' ||
		char == ':' || char == ';'
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
