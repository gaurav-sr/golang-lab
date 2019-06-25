package main

import (
	"fmt"
	"log"
)

type growmoney interface {
	calculate() float64
}

type lifepurpose struct {
	money float64
}

func (l lifepurpose) calculate() float64 {
	return l.money * 2.1
}

func main() {
	l := lifepurpose{money: 23.6}
	amount := l.calculate()
	fmt.Println(amount)
	log.Printf("Amount is %f", amount)
}
