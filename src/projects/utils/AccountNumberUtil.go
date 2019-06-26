package main

import (
	"fmt"
)

func main() {
	var accountNumber string = "26467164768113"
	masked := "xxxx" + accountNumber[4:len(accountNumber)]
	fmt.Println(masked)

}
