package main

import (
	"ASCII-Art/printart"
	"ASCII-Art/readandprocess"
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) < 1 || len(arguments) > 2 {
		fmt.Println(`Incorrect number of arguments.
Usage:
go run . "hello"
	OR
Usage: go run. "hello" "thinkertoy.txt"`)
		return
	}

	inputString := arguments[0]

	var bannerFont string

	// Determines the banner file to access if specified in the arguments else selects "standard.txt" as the default file
	if len(arguments) == 2 {
		bannerFont = arguments[1]
	} else {
		bannerFont = "standard.txt"
	}

	bannerFile, err := readandprocess.ReadAndProcess(bannerFont)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	printart.PrintArt(bannerFile, inputString)
}
