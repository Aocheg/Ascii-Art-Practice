package main

import (
	"fmt"
	"os"
	"strings"
)

func AsciiAllChar(text string) string {
	data, _ := os.ReadFile("standard.txt")
	words := string(data)

	SplitedWords := strings.Split(words, "\n")
	splittext := strings.Split(text, "\n")

	var result string

	for _, value := range splittext {
		for row := 0; row < 8; row++ {
			for _, char := range value {
				start := (int(char) - 32) * 9
				result += (SplitedWords[start+row])
			}
			result += "\n"

		}
		fmt.Println("\n")
	}

	return result

}
func main() {
	fmt.Println(AsciiAllChar("1Hello 2There"))
}
