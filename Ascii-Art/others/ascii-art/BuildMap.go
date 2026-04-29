package main

import "strings"

func BuildMap(content string) map[rune][]string {
	lines := strings.Split(content, "\n")

	asciiMap := make(map[rune][]string)

	char := 32

	for i := 0; i+8 < len(lines); i += 9 {
		asciiMap[rune(char)] = append([]string{}, lines[i:i+8]...)
		char++
	}
	return asciiMap
}
