package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//argument passing
	if len(os.Args) != 3 {
		return
	}
	//input validation
	var input string
	banner := "standard"

	if len(os.Args) == 2 {
		input = os.Args[1]
	} else if len(os.Args) == 3 {
		input = os.Args[1]
		banner = os.Args[2]
	} else {
		return
	}

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		fmt.Println("invalid banner")
		return
	}
	// fmt.Print(input)

	//newline handling
	input = strings.ReplaceAll(input, "\\n", "\n")
	// fmt.Println(input)

	//read banner file
	data, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println("error handling file")
		return
	}
	content := string(data)

	// fmt.Println(content[:200])

	lines := strings.Split(content, "\n")

	asciiMap := make(map[rune][]string)

	char := 32

	for i := 0; i+8 < len(lines); i += 9 {

		asciiMap[rune(char)] = append([]string{}, lines[i:i+8]...)
		char++

	}
	// fmt.Println(asciiMap['A'])
	// for _, line := range asciiMap['A'] {
	// 	fmt.Println(line)
	// }

	for _, word := range strings.Split(input, "\n") {

		for i := 0; i < 8; i++ {
			for _, ch := range word {
				fmt.Print(asciiMap[ch][i])
			}
			fmt.Println()
		}
	}

}
