package main

import "fmt"

func main() {
	count := 0
	largest := 0
	for count < 10 {
		fmt.Print("Enter a number: ")
		var number int
		fmt.Scanln(&number)
		count = count + 1
		if number > largest {
			largest = number
		}
	}
	fmt.Println(largest)
}
