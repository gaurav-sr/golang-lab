package main

import (
	"fmt"
	"log"
)

type GrowMoney interface {
	Calculate() float64
}

type lifepurpose struct {
	money float64
}

func (l lifepurpose) Calculate() float64 {
	return l.money * 2.1
}

func main() {
	l := lifepurpose{money: 23.6}
	amount := l.Calculate()
	fmt.Println(amount)
	log.Printf("Amount is %f", amount)
}
