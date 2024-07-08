package readandprocess

import (
	"ASCII-Art/checksum"
	"fmt"
	"os"
	"strings"
)

func ReadAndProcess(bannerFile string) ([]string, error) {
	bannerFileData, err := os.ReadFile(bannerFile)
	if err != nil {
		return nil, err
	}

	// Check file integrity
	fileHash := checksum.CheckFileValidity(bannerFileData)

	switch bannerFile {
	case "standard.txt":
		if fileHash != "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf" {
			fmt.Printf("Error: Possible file corruption with %v\n", bannerFile)
			os.Exit(1)
		}
	case "shadow.txt":
		if fileHash != "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73" {
			fmt.Printf("Error: Possible file corruption with %v\n", bannerFile)
			os.Exit(1)
		}
	case "thinkertoy.txt":
		if fileHash != "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3" {
			fmt.Printf("Error: Possible file corruption with %v\n", bannerFile)
			os.Exit(1)
		}
	}

	var splitBannerFileData []string

	if bannerFile == "thinkertoy.txt" {
		splitBannerFileData = strings.Split(string(bannerFileData), "\r\n")
	} else {
		splitBannerFileData = strings.Split(string(bannerFileData), "\n")
	}

	return splitBannerFileData, nil
}
