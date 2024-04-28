package main

import "fmt"

func main() {
	var number int

	for {
		fmt.Println("Enter number:")
		fmt.Scanln(&number)

		if number == 1 || number == 2 {
			fmt.Println("Successful")
			break
		} else {
			fmt.Println("Invalid input. Please enter either 1 or 2.")
		}
	}
}
