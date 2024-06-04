package readandprocess

import (
	"fmt"
	"os"
	"strings"
)

func ReadAndProcess(bannerFile string) ([]string, error) {
	bannerFileData, err := os.ReadFile(bannerFile)
	if err != nil {
		return nil, err
	}

	var splitBannerFileData []string

	if bannerFile == "thinkertoy.txt" {
		splitBannerFileData = strings.Split(string(bannerFileData), "\r\n")
	} else {
		splitBannerFileData = strings.Split(string(bannerFileData), "\n")
	}

	if len(splitBannerFileData) != 856 {
		return nil, fmt.Errorf("corrupted file")
	}

	return splitBannerFileData, nil
}
