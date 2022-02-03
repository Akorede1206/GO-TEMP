package main

import "testing"

func TestMySum(t *testing.T) { //Pointer of T assigned to varriable t

	type test struct {
		data   []int
		answer int
	}

	tests := []test{ //composite literal - slice of test
		test{[]int{21, 21}, 42},
		test{[]int{3, 50, 21}, 74}, //values Composite literal - (slice of ints (values))
		test{[]int{1, 1}, 2},
		test{[]int{-1, 0, 1}, 0},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer { //mySum Function writting in main.go
			t.Error("Expected", v.answer, "Got", x) //Assertion
		}
	}
}
