package main

import "fmt"

func main() {
	var length int
	fmt.Println("Enter length of triangle (between 1 to 10): ")
	fmt.Scanln(&length)

	if length < 1 || length > 10 {
		fmt.Println("Invalid length. Please enter a length between 1 and 10.")
		return
	}

	for index := 1; index <= length; index++ {
		for idx := 1; idx <= index; idx++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
