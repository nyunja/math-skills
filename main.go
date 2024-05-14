package main

import (
	"os"
	"runtime"

	"math-skills/functions"
)

func main() {
	// Exit the program if the arguments are less or more than expected
	if len(os.Args) != 2 {
		os.Stdout.WriteString("Error: please provide the required arguments\n")
		return
	}

	// Check filename provided and exit if the name is invalid
	fileName := os.Args[1]
	if fileName != "data.txt" {
		os.Stdout.WriteString("Error: please provide the correct file\n")
		return
	}

	// Read the contents of the file
	content, err := os.ReadFile(fileName)
	if err != nil {
		os.Stdout.WriteString("Error: unable to read " + fileName + " content\n")
		return
	}

	// Checks whether data.txt is empty
	if len(content) == 0 {
		os.Stdout.WriteString("Error: " + fileName + " is empty\n")
		return
	}

	var numStr []string

	// Apply split pattern depending on operating system
	ops := runtime.GOOS
	switch ops {
	case "windows":
		numStr = functions.SplitString(string(content), "\r\n")
	case "linux":
		numStr = functions.SplitString(string(content), "\n")
	}

	// Call function to handle calculation
	functions.HandleCalculation(numStr)
}
