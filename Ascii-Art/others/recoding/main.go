package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("error")
	}
	fmt.Println(LoadBanner(os.Args[1]))
}
