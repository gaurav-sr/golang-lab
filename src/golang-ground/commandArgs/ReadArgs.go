package main

import (
	"fmt"
	"os"
)

func readArgs() {
	fmt.Printf("This program accepts user input and prints number of characters\n")
	arguments := os.Args
	printArgsFrom(arguments, 0)
}

func printArgsFrom(arguments []string, position int) {
	for idx, arg := range arguments[position:] {
		fmt.Printf("Argument at %d is %s and #characters=%d\n", idx, arg, len(arg))
	}
}
