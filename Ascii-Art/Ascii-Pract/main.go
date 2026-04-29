package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

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
		fmt.Println("invalid banner:", banner, "\n", "Use: standard, shadow, or thinkertoy")
		return
	}

	input = strings.ReplaceAll(input, "\\n", "\n")

	content := LoadBanner(banner)
	asciiMap := BuildMap(content)
	RenderAscii(input, asciiMap)

}
