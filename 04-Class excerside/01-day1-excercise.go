package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)

		if i%3 == 0 {
			fmt.Println("\n\nFizz: Its divsible by 3")

			{
				if i%5 == 0 {
					fmt.Println("\t\t\tBuzz:Its divsible by 5")

					{
						if i%3 == 0 && i%5 == 0 {
							fmt.Println("\t\t\t\tFizzBuzz: Its divsible by 3 & 5")
						}
					}
				}
			}
		}

	}

}
