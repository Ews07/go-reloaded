package main

import (
	"fmt"
	"os"

	"functions"
)

func main() {
	args := os.Args[1:]
	//Check if if there two arguments
	if len(args) != 2 {
		fmt.Println("Usage: go run . <inputfile> <outputfile>")
		return
	}

	inputFilename := args[0]
	outputFilename := args[1]
	//Read the content from input file
	inputText, err := functions.ReadFromFile(inputFilename)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	//Call the functions that modify the input text
	inputText = functions.FlagsManipulation(inputText)
	//Write to the output file
	err = functions.WriteToFile(outputFilename, inputText)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
	}
}
