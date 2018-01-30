package main

import "fmt"

/*The variable is initialized before the init() and main() functions are called*/
var year = 2018

/**
The init() function is called after the variable year is initialized and
before the main function is called. This function can be used to initialize variables.
*/
func init() {
	fmt.Println("Current year ", year)
	year = year + 1
}

/**
This is GO main entry point and called after the package level variables are initialized
and after init() function is called.
*/
func main() {
	fmt.Println("Next Year ", year)
}
