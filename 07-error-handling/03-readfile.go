package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("names.txt") //implementing writer inetrface
	if err != nil {
		fmt.Println("err happened", err) //Printing an error
		log.Println("err happened", err) //logging an error
	}

}
