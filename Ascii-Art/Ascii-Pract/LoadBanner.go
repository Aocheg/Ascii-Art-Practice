package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(banner string) string {

	data, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return string(data)
}

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

func RenderAscii(input string, asciiMap map[rune][]string) {
	if input == "" {
		return
	}

	if input == "\n" {
		fmt.Println()
		return
	}

	for _, word := range strings.Split(input, "\n") {
		if word == "" {
			fmt.Println()
			continue
		}

		for row := 0; row < 8; row++ {
			for _, ch := range word {
				val, exists := asciiMap[ch]

				if !exists {
					fmt.Println("Invalid Character:", string(ch))
					return
				}
				fmt.Print(val[row])
			}
			fmt.Println()
		}
	}
}
