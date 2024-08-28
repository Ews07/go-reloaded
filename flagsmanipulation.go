package functions

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func FlagsManipulation(input string) string {
	// Split the input text line by line
	lines := strings.Split(input, "\n")
	for k, line := range lines {
		// Split the line word by word without spces
		words := strings.Fields(line)

		for i := 0; i < len(words); i++ {
			if words[i] == "(up," {
				// Check if the flag is correctly write
				if i+1 < len(words) && strings.Contains(words[i+1], ")") {
					numStr := strings.TrimSuffix(words[i+1], ")")
					num, err := strconv.Atoi(numStr)
					if err != nil {
						fmt.Println("invalid num")
						i += 2
						continue
					}
					// initialize index to apply the flag for the words Desirable
					startIndex := i - num
					if startIndex < 0 {
						startIndex = 0
					}

					for j := startIndex; j < i; j++ {
						words[j] = strings.ToUpper(words[j])
					}
					// append the words without the flag and his argument
					words = append(words[:i], words[i+2:]...)
					i--
				}
			} else if words[i] == "(up)" {
				if i > 0 {
					words[i-1] = strings.ToUpper(words[i-1])
				}
				// append the words without the flag
				words = append(words[:i], words[i+1:]...)
				i--
			} else if words[i] == "(low," {
				// Check if the flag is correctly write
				if i+1 < len(words) && strings.Contains(words[i+1], ")") {
					numStr := strings.TrimSuffix(words[i+1], ")")
					num, err := strconv.Atoi(numStr)
					if err != nil {
						fmt.Println("invalid num")
						i += 2
						continue
					}
					// initialize index to apply the flag for the words Desirable
					startIndex := i - num
					if startIndex < 0 {
						startIndex = 0
					}

					for j := startIndex; j < i; j++ {
						words[j] = strings.ToLower(words[j])
					}
					// append the words without the flag and his argument
					words = append(words[:i], words[i+2:]...)
					i--
				}
			} else if words[i] == "(low)" {
				if i > 0 {
					words[i-1] = strings.ToLower(words[i-1])
				}
				// append the words without the flag
				words = append(words[:i], words[i+1:]...)
				i--
			} else if words[i] == "(cap," {
				// Check if the flag is correctly write
				if i+1 < len(words) && strings.Contains(words[i+1], ")") {
					numStr := strings.TrimSuffix(words[i+1], ")")
					num, err := strconv.Atoi(numStr)
					if err != nil {
						fmt.Println("invalid num")
						i += 2
						continue
					}

					startIndex := i - num
					if startIndex < 0 {
						startIndex = 0
					}
					// initialize index to apply the flag for the words Desirable
					for j := startIndex; j < i; j++ {
						firstRune, size := utf8.DecodeRuneInString(words[j])
						words[j] = strings.ToUpper(string(firstRune)) + strings.ToLower(words[j][size:])
					}
					// append the words without the flag and his argument
					words = append(words[:i], words[i+2:]...)
					i--
				}
			} else if words[i] == "(cap)" {
				if i > 0 {
					firstRune, size := utf8.DecodeRuneInString(words[i-1])
					words[i-1] = strings.ToUpper(string(firstRune)) + strings.ToLower(words[i-1][size:])
				}
				words = append(words[:i], words[i+1:]...)
				i--
			} else if words[i] == "(hex)" {
				if i > 0 && IsHexadecimal(words[i-1]) {
					decimalValue, err := strconv.ParseInt(words[i-1], 16, 64)
					if err == nil {
						words[i-1] = fmt.Sprintf("%d", decimalValue)
					}
				}
				words = append(words[:i], words[i+1:]...)
				i--
			} else if words[i] == "(bin)" {
				if i > 0 && IsBinary(words[i-1]) {
					decimalValue, err := strconv.ParseInt(words[i-1], 2, 64)
					if err == nil {
						words[i-1] = fmt.Sprintf("%d", decimalValue)
					}
				}
				words = append(words[:i], words[i+1:]...)
				i--
			} else {
				// add "n" to a if there a vowel followed
				if (words[i] == "a" || words[i] == "A") && (i+1 < len(words) && IsVowel(words[i+1])) {
					words[i] = words[i] + "n"
				}
			}
		}
		lines[k] = strings.Join(words, " ")
		//applied punctuations and single qote functions of each line
		lines[k] = HandlePunctuations(lines[k])
		lines[k] = SingleQuoteHandler(lines[k])
		lines[k] = strings.TrimSuffix(lines[k], " ")
	}
	return strings.Join(lines, "\n")
}
