package functions

import (
	"strings"
	"unicode"
)

func SingleQuoteHandler(words string) string {
	isFirst := true
	result := ""
	skip := false
	punc := ".,;:!?"
	for i, char := range words {
		// Not add a space if there a puncuation after the quote
		if strings.ContainsRune(punc, char) && (result != "" && result[len(result)-1] == ' ') {
			result = result[:len(result)-1]
		}
		if char == '\'' {
			// Check the boundaries and adjacent characters
			if (i == 0 || !unicode.IsLetter(rune(words[i-1]))) || (i == len(words)-1 || !unicode.IsLetter(rune(words[i+1]))) {
				if isFirst {
					// Add a space before the quote if necessary
					if len(result) > 0 && result[len(result)-1] != ' ' {
						result += " "
					}
					result += "'"
					isFirst = false
					// Skip the next space if any
					if i+1 < len(words) && words[i+1] == ' ' {
						skip = true
					}
				} else {
					// Remove the trailing space before adding the quote if necessary
					if len(result) > 0 && result[len(result)-1] == ' ' {
						result = result[:len(result)-1]
					}
					result += "'"
					isFirst = true
					// Add a space after the quote if necessary
					if i+1 < len(words) && words[i+1] != ' ' {
						result += " "
					}
				}
				continue
			}
		} else {
			if skip {
				skip = false
				continue
			}
		}
		result += string(char)
	}

	return result
}
