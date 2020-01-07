package main

import "testing"

func Even(i int) bool {		// Exported function
	return i%2 == 0
}

func Odd(i int) bool {		// Exported function
	return i%2 != 0
}

func TestEvenOdd(t *testing.T) {
	if !Even(10){
		t.Log("10 must be even")
		t.Fail()
	}
	if Even(7){
		t.Log("7 can't be even")
		t.Fail()
	}
}
