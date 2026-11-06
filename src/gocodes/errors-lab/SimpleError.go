package main

import (
	"errors"
	"fmt"
)

func main() {
	p := player{name: "Gaurav", age: 32}
	code, msg := p.checkEligibility()
	if msg != nil {
		fmt.Printf("Player %s. Message : %s. Code : %d", p.name, msg, code)
	} else {
		fmt.Printf("Player %s eligible", p.name)
	}
	p = player{name: "Aarush", age: 10}
	code, msg = p.checkEligibility()
	if msg != nil {
		fmt.Printf("\nPlayer %s. Message : %s. Code : %d", p.name, msg, code)
	}
}

type player struct {
	name string
	age  int
}

func (p *player) checkEligibility() (int, error) {
	if p.age < 18 {
		return -1, errors.New("Ineligible age")
	}
	return 0, nil
}
