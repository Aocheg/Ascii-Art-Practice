package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func Ascii(text rune) string {
// 	data, _ := os.ReadFile("standard.txt")
// 	words := string(data)

// 	SplitedWords := strings.Split(words, "\n")

// 	start := (int(text) - 32) * 9
// 	var result string

// 	for row := 0; row < 8; row++ {
// 		result += (SplitedWords[start+row]) + "\n"
// 	}
// 	return result

// }
// // func main() {
// // 	fmt.Println(Ascii('S'))
// // }
