package main

import (
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	font := make(map[rune][]string)

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return font, err
	}
	lines := strings.Split(string(data), "\n")

	for char := ' '; char <= '~'; char++ {
		start := (int(char) - 32) * 9
		font[char] = lines[start+1 : start+9]
	}
	return font, nil
}
