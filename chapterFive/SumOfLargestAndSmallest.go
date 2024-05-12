package main

import (
	"fmt"
	"math"
)

func main() {
	count := 0
	largest := math.MinInt64
	smallest := math.MaxInt64
	sum := 0
	for count < 10 {
		fmt.Print("Enter number : ")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			return
		}
		if number == -1 {
			break
		}
		if number > largest {
			largest = number
		}
		if number < smallest {
			smallest = number
		}
		count++
	}
	if count < 2 {
		fmt.Println("You need to enter at least two numbers.")
		return
	}
	sum = largest + smallest
	fmt.Printf("The sum of the largest and smallest numbers is %d\n", sum)
}
