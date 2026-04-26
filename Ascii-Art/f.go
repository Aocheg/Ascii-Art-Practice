package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func AsciArt(words string, banner string) string {
// 	data, _ := os.ReadFile("standard.txt")
// 	text := string(data)

// 	Splittext := strings.Split(text, "\n")
// 	splitwords := strings.Split(words, "\n")

// 	var result string

// 	for _, val := range splitwords {
// 		for row := 0; row < 8; row++ {
// 			for _, char := range val {
// 				start := (int(char) - 32) * 9
// 				result += (Splittext[start+row])
// 			}
// 			result += "\n"
// 		}
// 		fmt.Println("\n")
// 	}
// 	return result
// }

// // func main() {
// // 	fmt.Println(AsciArt("k"))
// // }
