package printart

import (
	"fmt"
	"strings"
)

func PrintArt(bannerFileSlice []string, inputString string) {
	if inputString == "\\n" {
		fmt.Println()
		return
	} else if inputString == "" {
		return
	} else if inputString == "\\t" {
		fmt.Println("    ")
		return
	}

	// Handle unprintable sequences
	unprintableSequences := []string{"\\a", "\\b", "\\v", "\\f", "\\r"}

	for _, unprintable := range unprintableSequences {
		if strings.Contains(inputString, unprintable) {
			fmt.Println("Input string contains an unprintable sequence")
			return
		}
	}

	tabCharText := strings.ReplaceAll(inputString, "\\t", "    ")
	newlineCharText := strings.ReplaceAll(tabCharText, "\\n", "\n")

	splitArguments := strings.Split(newlineCharText, "\\n")

	// Handle foreign input strings
	for _, splitArg := range splitArguments {
		for _, char := range splitArg {
			if char < 32 || char > 126 {
				fmt.Println("Input string contains an unprintable character")
				return
			}
		}
	}

	for _, text := range splitArguments {
		if text == "" {
			fmt.Println()
			continue
		}

		for j := 0; j < 8; j++ {
			for _, char := range text {
				startingIndex := int(char-32)*9 + 1
				fmt.Printf(bannerFileSlice[startingIndex+j])
			}
			fmt.Println()
		}
	}
}
