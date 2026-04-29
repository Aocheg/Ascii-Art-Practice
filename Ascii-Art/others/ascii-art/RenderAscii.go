package main

import (
	"fmt"
	"strings"
)

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
