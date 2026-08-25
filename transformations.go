package main

import (
	"fmt"
	"strconv"
	"strings"
)

func processWords(words []string) ([]string, error) {
	result := []string{}

	for i := 0; i < len(words); i++ {
		word := strings.ToLower(words[i])

		// Hexadecimal
		if word == "(hex)" {
			if len(result) == 0 {
				return nil, fmt.Errorf("(hex) has no previous word")
			}

			number, err := strconv.ParseInt(result[len(result)-1], 16, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid hexadecimal number: %s", result[len(result)-1])
			}

			result[len(result)-1] = strconv.FormatInt(number, 10)
			continue
		}

		// Binary
		if word == "(bin)" {
			if len(result) == 0 {
				return nil, fmt.Errorf("(bin) has no previous word")
			}

			number, err := strconv.ParseInt(result[len(result)-1], 2, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid binary number: %s", result[len(result)-1])
			}

			result[len(result)-1] = strconv.FormatInt(number, 10)
			continue
		}

		// Uppercase
		if word == "(up)" {
			if len(result) == 0 {
				return nil, fmt.Errorf("(up) has no previous word")
			}

			if !hasLetter(result[len(result)-1]) {
				return nil, fmt.Errorf("(up) cannot be applied to: %s", result[len(result)-1])
			}

			result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			continue
		}

		// Lowercase
		if word == "(low)" {
			if len(result) == 0 {
				return nil, fmt.Errorf("(low) has no previous word")
			}

			if !hasLetter(result[len(result)-1]) {
				return nil, fmt.Errorf("(low) cannot be applied to: %s", result[len(result)-1])
			}

			result[len(result)-1] = strings.ToLower(result[len(result)-1])
			continue
		}

		// Capitalize
		if word == "(cap)" {
			if len(result) == 0 {
				return nil, fmt.Errorf("(cap) has no previous word")
			}

			if !hasLetter(result[len(result)-1]) {
				return nil, fmt.Errorf("(cap) cannot be applied to: %s", result[len(result)-1])
			}

			result[len(result)-1] = capitalize(result[len(result)-1])
			continue
		}

		// (up, N)
		if word == "(up," {
			if i+1 >= len(words) {
				return nil, fmt.Errorf("(up, N) is missing its number")
			}

			numberText := strings.TrimSuffix(words[i+1], ")")
			number, err := strconv.Atoi(numberText)

			if err != nil {
				return nil, fmt.Errorf("invalid number in (up, N)")
			}

			if number > len(result) {
				return nil, fmt.Errorf("(up, %d) has too many words", number)
			}

			for j := len(result) - number; j < len(result); j++ {
				if !hasLetter(result[j]) {
					return nil, fmt.Errorf("(up) cannot be applied to: %s", result[j])
				}

				result[j] = strings.ToUpper(result[j])
			}

			i++
			continue
		}

		// (low, N)
		if word == "(low," {
			if i+1 >= len(words) {
				return nil, fmt.Errorf("(low, N) is missing its number")
			}

			numberText := strings.TrimSuffix(words[i+1], ")")
			number, err := strconv.Atoi(numberText)

			if err != nil {
				return nil, fmt.Errorf("invalid number in (low, N)")
			}

			if number > len(result) {
				return nil, fmt.Errorf("(low, %d) has too many words", number)
			}

			for j := len(result) - number; j < len(result); j++ {
				if !hasLetter(result[j]) {
					return nil, fmt.Errorf("(low) cannot be applied to: %s", result[j])
				}

				result[j] = strings.ToLower(result[j])
			}

			i++
			continue
		}

		// (cap, N)
		if word == "(cap," {
			if i+1 >= len(words) {
				return nil, fmt.Errorf("(cap, N) is missing its number")
			}

			numberText := strings.TrimSuffix(words[i+1], ")")
			number, err := strconv.Atoi(numberText)

			if err != nil {
				return nil, fmt.Errorf("invalid number in (cap, N)")
			}

			if number > len(result) {
				return nil, fmt.Errorf("(cap, %d) has too many words", number)
			}

			for j := len(result) - number; j < len(result); j++ {
				if !hasLetter(result[j]) {
					return nil, fmt.Errorf("(cap) cannot be applied to: %s", result[j])
				}

				result[j] = capitalize(result[j])
			}

			i++
			continue
		}

		// Normal word
		result = append(result, words[i])
	}

	return result, nil
}
