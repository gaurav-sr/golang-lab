package main

import "os"
import "strconv"
import "fmt"

/**
This program add the arguments and prints the result. If the argument is not an integer , it throws error.
**/
func sumArgs() {
	fmt.Printf("This program add the arguments and prints the result. If the argument is not an integer , it throws error\n")
	args := os.Args
	var sum int
	for _, arg := range args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Only integer argument allowed")
			panic(err)
		}
		sum = sum + num
	}
	fmt.Println("Sum is ", sum)
}
