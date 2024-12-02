package utils

import (
	"os"
	"strings"
)

func Parse(fileName string) ([]string, error) {
	rawFile, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	// split lines
	lines := strings.Split(string(rawFile), "\n")
	// get last index
	lastIndex := len(lines) - 1
	// remove empty blank line
	lines = lines[:lastIndex]

	return lines, nil
}
