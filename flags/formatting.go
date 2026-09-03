package flags

import "strings"

func FormatPunctuation(words []string) string {
	result := ""
	inQuote := false

	for _, word := range words {

		// Standalone quote
		if word == "'" {
			if !inQuote { //if NOT already inside a quote
				if result != "" { //add space
					result += " "
				}
				result += "'"
				inQuote = true
			} else {
				result = strings.TrimSpace(result)
				result += "'"
				inQuote = false
			}
			continue
		}

		// Word starts with a quote
		if strings.HasPrefix(word, "'") {
			if result != "" {
				result += " "
			}
			result += word
			inQuote = true
			continue
		}

		// Word ends with a quote
		if strings.HasSuffix(word, "'") {
			if result != "" && !strings.HasSuffix(result, "'") {
				result += " "
			}
			result += word
			inQuote = false
			continue
		}

		// Closing parenthesis
		if word == ")" {
			result = strings.TrimSpace(result)
			result += ")"
			continue
		}

		if strings.HasPrefix(word, ")") {
			result = strings.TrimSpace(result)
			result += word
			continue
		}
		// Check if the entire word is punctuation
		allPunctuation := true

		for _, char := range word {
			if !isPunctuation(char) {
				allPunctuation = false
				break
			}
		}

		// Keep punctuation groups together: ..., !?, !!!
		if allPunctuation {
			result += word
			continue
		}

		// Punctuation at the beginning of a word
		if len(word) > 1 && isPunctuation(rune(word[0])) {
			result += string(word[0])
			result += " " + word[1:]
			continue
		}

		// Normal word
		if result != "" && result != "(" {
			if !strings.HasSuffix(result, "'") || !inQuote {
				result += " "
			}
		}
		result += word
	}
	return result
}
