package main

import (
	"bytes"
	"os"
	"testing"
)

// func TestRenderAscii(t *testing.T) {
// 	asciiMap := map[rune][]string{
// 		'A': {
// 			" A ",
// 			" A A ",
// 			"AAAAA",
// 			"A   A",
// 			"A   A",
// 			"A   A",
// 			"A   A",
// 			"A   A",
// 		},
// 	}

// 	input := "A"

// 	old := os.Stdout
// 	r, w, _ := os.Pipe()
// 	os.Stdout = w

// 	RenderAscii(input, asciiMap)

// 	w.Close()
// 	os.Stdout = old

// 	var result bytes.Buffer
// 	result.ReadFrom(r)

// 	expected :=

// 		` A
// 	A A
// 	AAAAA
// 	A   A
// 	A   A
// 	A   A
// 	A   A
// 	A   A
// 	`
// 	if result.String() != expected {
// 		t.Errorf("expected:\n%v\ngot:\n%v", expected, result.String())
// 	}
// }

func TestRenderAscii(t *testing.T) {

	asciiMap := map[rune][]string{
		'A': {
			"  A  ",
			" A A ",
			"AAAAA",
			"A   A",
			"A   A",
			"A   A",
			"A   A",
			"A   A",
		},
	}

	input := "A"

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	RenderAscii(input, asciiMap)

	w.Close()
	os.Stdout = old

	var result bytes.Buffer
	result.ReadFrom(r)

	expected :=
		`  A  
 A A 
AAAAA
A   A
A   A
A   A
A   A
A   A
`

	if result.String() != expected {
		t.Errorf("expected:\n%v\ngot:\n%v", expected, result.String())
	}
}
