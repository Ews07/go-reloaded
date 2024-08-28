package functions

import "regexp"

func HandlePunctuations(input string) string {
	punct := regexp.MustCompile(` *[.,!?;: ]+ *`)
	//replacing punct in input with the string who return this function
	output := punct.ReplaceAllStringFunc(input, func(s string) string {
		result := ""
		for _, letter := range s {
			if letter != ' ' {
				result += string(letter)
			}
		}
		return result + " "
	})
	return output
}
