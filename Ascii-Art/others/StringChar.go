package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func AsciiArt(text string) string {
// 	data, _ := os.ReadFile("standard.txt")
// 	words := string(data)

// 	SplitedWords := strings.Split(words, "\n")

// 	var result string

// 	for row := 0; row < 8; row++ {
// 		for _, char := range text {
// 			start := (int(char) - 32) * 9
// 			result += (SplitedWords[start+row])
// 		}
// 		result += "\n"

// 	}
// 	return result

// }

// func main() {
// 	fmt.Println(AsciiArt("Gabriel\n"))
// }
