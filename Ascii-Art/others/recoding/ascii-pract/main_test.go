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
// 		`  A
// A A
// AAAAA
// A   A
// A   A
// A   A
// A   A
// A   A
// `
// 	if result.String() != expected {
// 		t.Errorf("expected:\n%v\ngot:\n%v", expected, result.String())
// 	}
// }

// func TestRenderAscii(t *testing.T) {

// 	asciiMap := map[rune][]string{
// 		'A': {
// 			"  A  ",
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
// 		`  A  
//  A A 
// AAAAA
// A   A
// A   A
// A   A
// A   A
// A   A
// `

// 	if result.String() != expected {
// 		t.Errorf("expected:\n%v\ngot:\n%v", expected, result.String())
// 	}
// }


package main

import (
	"bytes"
	"os"
	"testing"
)

func captureOutput(f func()) string {

	old := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	f()

	w.Close()

	os.Stdout = old

	var buf bytes.Buffer

	buf.ReadFrom(r)

	return buf.String()
}

func TestRenderSingleCharacter(t *testing.T) {

	asciiMap := map[rune][]string{
		'A': {
			"A",
			"A",
			"A",
			"A",
			"A",
			"A",
			"A",
			"A",
		},
	}

	output := captureOutput(func() {
		RenderAscii("A", asciiMap)
	})

	expected :=
		`A
A
A
A
A
A
A
A
`

	if output != expected {
		t.Errorf("expected:\n%v\ngot:\n%v", expected, output)
	}
}

func TestRenderMultipleCharacters(t *testing.T) {

	asciiMap := map[rune][]string{
		'A': {
			"A", "A", "A", "A",
			"A", "A", "A", "A",
		},
		'B': {
			"B", "B", "B", "B",
			"B", "B", "B", "B",
		},
	}

	output := captureOutput(func() {
		RenderAscii("AB", asciiMap)
	})

	expected :=
		`AB
AB
AB
AB
AB
AB
AB
AB
`

	if output != expected {
		t.Errorf("expected:\n%v\ngot:\n%v", expected, output)
	}
}

func TestEmptyInput(t *testing.T) {

	asciiMap := map[rune][]string{}

	output := captureOutput(func() {
		RenderAscii("", asciiMap)
	})

	if output != "" {
		t.Errorf("expected empty output, got %q", output)
	}
}

func TestNewlineInput(t *testing.T) {

	asciiMap := map[rune][]string{}

	output := captureOutput(func() {
		RenderAscii("\n", asciiMap)
	})

	expected := "\n"

	if output != expected {
		t.Errorf("expected newline, got %q", output)
	}
}

func TestDoubleNewline(t *testing.T) {

	asciiMap := map[rune][]string{
		'A': {
			"A", "A", "A", "A",
			"A", "A", "A", "A",
		},
	}

	output := captureOutput(func() {
		RenderAscii("A\n\nA", asciiMap)
	})

	expected :=
		`A
A
A
A
A
A
A
A

A
A
A
A
A
A
A
A
`

	if output != expected {
		t.Errorf("expected:\n%v\ngot:\n%v", expected, output)
	}
}

func TestInvalidCharacter(t *testing.T) {

	asciiMap := map[rune][]string{
		'A': {
			"A", "A", "A", "A",
			"A", "A", "A", "A",
		},
	}

	output := captureOutput(func() {
		RenderAscii("😊", asciiMap)
	})

	expected := "Unsupported character: \"😊\"\n"

	if output != expected {
		t.Errorf("expected %q, got %q", expected, output)
	}
}

func TestLoadBanner(t *testing.T) {

	content := loadBanner("standard")

	if content == "" {
		t.Errorf("expected banner content, got empty string")
	}
}

func TestBuildMap(t *testing.T) {

	content := loadBanner("standard")

	asciiMap := buildMap(content)

	if len(asciiMap) == 0 {
		t.Errorf("asciiMap is empty")
	}

	if _, exists := asciiMap['A']; !exists {
		t.Errorf("character A missing from map")
	}
}