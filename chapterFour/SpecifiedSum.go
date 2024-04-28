package main

import (
	"fmt"
)

func main() {
	var number, number1, sum int
	number1 = 10

	for {
		fmt.Print("Enter number: ")
		fmt.Scanln(&number)

		if number == number1 || number > number1 {
			fmt.Println("Successful")
			sum += number
			break
		}
	}

	fmt.Println("Sum:", sum)
}
