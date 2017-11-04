package main

import "os"
import "strconv"
import "fmt"

func sumArgs() {

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
