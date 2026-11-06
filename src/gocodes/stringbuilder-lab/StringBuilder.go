package main

import (
	"strings"
)

func main() {
	var sb strings.Builder
	sb.Grow(128)

	sb.WriteString("Hello")
	sb.WriteString(" ")
	sb.WriteString("World")
	println(sb.String())

	// convert to upper and lower case
	sb.Reset()
	sb.WriteString("A quick brown fox")
	println(strings.ToUpper(sb.String()))
	println(strings.ToLower(sb.String()))

	//Split into words
	sb.Reset()
	sb.WriteString("Miles to go before I sleep")
	words := strings.Fields(sb.String())
	for _, word := range words {
		println(word)
	}

	//Join words
	joined := strings.Join(words, "-")
	println(joined)
}
