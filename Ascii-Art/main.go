package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println(`Invalid Argument. Usage: go run . <text> <banner>`)
		return
	}
	input := os.Args[1]
	banner := "standard"

	if len(os.Args) == 3 {
		banner = os.Args[2]
	}

	input = strings.ReplaceAll(input, "\\n", "\n")

	result := AsciiAllChar(input, banner)

	fmt.Print(result)
}
