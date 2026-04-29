package main

import (
	"fmt"
	"os"
)

func LoadBanner(banner string) string {
	data, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return string(data)
}
