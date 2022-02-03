package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	changeMe(
		person{
			"ade",
			"ade",
		},
	)

	changeMe(
		person{
			"Tade",
			"Tade",
		},
	)

	p1 := person{
		"pade",
		"Pade",
	}

	p2 := person{
		"pade",
		"Pade",
	}

	people := []person{p1, p2}

	fmt.Println(people)
}

func changeMe(p person) {
	fmt.Println(p)
	p = person{"Badmus",
		"Bade"}
	fmt.Println(p)
	fmt.Println(p.first)
	fmt.Println(p.last)
}

bs, err : json.
