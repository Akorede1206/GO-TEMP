package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
}

func main() {
	s := `[{"First":"pade","Last":"Pade"},{"First":"pade","Last":"Pade"}]`
	bs := []byte(s) // slice of bites
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	people := []person{}
	//var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all of the data", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last)
	}
}
