package main

import (
	"fmt"
)

func main() {
	var number float64
	fmt.Println("Enter number: ")
	_, err := fmt.Scanln(&number)
	if err != nil {
		return
	}
	result := value(int(number))
	fmt.Println(result)
}

func factorial(num int) float64 {
	total := 1.0

	for input := num; input > 0; input-- {
		total *= float64(input)
	}
	return total
}

func value(num1 int) float64 {
	var value float64
	for index := 0; index < num1; index++ {
		value += 1 / factorial(index)
	}
	return value
}
