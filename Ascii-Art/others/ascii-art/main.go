package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	
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

	//banner validation
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		fmt.Println("invalid banner:", banner, "\n", "Use: standard, shadow, or thinkertoy")
		return
	}

	//newline handling
	input = strings.ReplaceAll(input, "\\n", "\n")

	//read banner file
	content := LoadBanner(banner)
	asciiMap := BuildMap(content)
	RenderAscii(input, asciiMap)
}
