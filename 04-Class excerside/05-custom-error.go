package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes", "never say never"},
	}

	toJSON(p1)
}

//toJSON needs to return an erro also
func toJSON(p1 person) error {

	_, err := json.Marshal(p1)

	fmt.Println(p1)
	if err != nil {
		fmt.Errorf("userlast %v                                                                                                                                                        (first %p) not found", p1.First, &p1.Last)
		//fmt.Println(err.Error())

	}
	return (err)
}
