package main

import (
	"os"
	"runtime"

	"math-skills/functions"
)

func main() {
	if len(os.Args) != 2 {
		os.Stdout.WriteString("Error: please provide the required arguments\n")
		return
	}
	fileName := os.Args[1]
	if fileName != "data.txt" {
		os.Stdout.WriteString("Error: please provide the correct file\n")
		return
	}
	content, err := os.ReadFile(fileName)
	if err != nil {
		os.Stdout.WriteString("Error: unable to read " + fileName + " content\n")
		return
	}
	if len(content) == 0 {
		os.Stdout.WriteString("Error: " + fileName + " is empty\n")
		return
	}
	var numStr []string

	ops := runtime.GOOS
	switch ops {
	case "windows":
		numStr = functions.SplitString(string(content), "\r\n")
	case "linux":
		numStr = functions.SplitString(string(content), "\n")
	}
	functions.HandleCalculation(numStr)
}
