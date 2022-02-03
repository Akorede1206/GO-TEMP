package main

import "testing"

func TestMySum(t *testing.T) { //Pointer of T assigned to varriable t
	x := mySum(2, 3)
	if x != 5 { //mySum Function writting in main.go
		t.Error("Expected", 5, "Got", x) //Assertion
	}
}
