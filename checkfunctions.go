package functions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func IsHexadecimal(s string) bool {
	match, _ := regexp.MatchString(`^[0-9A-Fa-f]+$`, s)

	if !match {
		fmt.Println("invalid hexadecimal number")
	}
	return match
}

func IsBinary(s string) bool {
	_, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println("Invalid binary number")
	}
	return err == nil
}

func IsVowel(word string) bool {
	boolean := false
	if len(word) == 0 {
		boolean = false
	}
	firstChar := strings.ToLower(word[0:1])
	if firstChar == "a" || firstChar == "e" || firstChar == "i" || firstChar == "o" || firstChar == "u" || firstChar == "h" {
		boolean = true
	}
	return boolean
}
