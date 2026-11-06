package main

import "testing"

func TestGreeting(t *testing.T) {
	greeting := GetGreeting()
	if greeting != GREETING {
		t.Errorf("Expected %s , Actual %s", GREETING, greeting)
	}
}
