package functions

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func IsvalidExtension(filename string) bool {
	return strings.HasSuffix(filename, ".txt")
}

func ReadFromFile(filename string) (string, error) {
	if !IsvalidExtension(filename) {
		return "", errors.New("file must be a txt file")
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strings.Join(lines, "\n"), nil
}

func WriteToFile(filename string, content string) error {
	if !IsvalidExtension(filename) {
		return errors.New("file must be a txt file")
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
