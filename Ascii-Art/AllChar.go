package main

import (
	"os"
	"strings"
)

func AsciiAllChar(text string, banner string) string {

	if text == "" {
		return ""
	}
	if text == "\n" {
		return "\n"
	}
	data, _ := os.ReadFile(banner + ".txt")
	words := string(data)

	SplitedWords := strings.Split(words, "\n")
	splittext := strings.Split(text, "\n")

	var result string

	for _, value := range splittext {
		if value == "" {
			if value == "" {
				result += "\n"
			}
			continue

		}

		for row := 0; row < 8; row++ {
			for _, char := range value {
				start := (int(char) - 32) * 9
				result += SplitedWords[start+row]

			}
			result += "\n"

		}
		// fmt.Println("\n")
	}

	return result

}

// func main() {
// 	fmt.Println(AsciiAllChar("1Hello 2There"))
// }
