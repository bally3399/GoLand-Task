package main

import "fmt"

func main() {
	var number int
	var number1 int
	fmt.Println("Enter number first number: ")
	_, err := fmt.Scanln(&number)
	if err != nil {
		return
	}
	fmt.Println("Enter second number: ")
	_, err = fmt.Scanln(&number1)
	if err != nil {
		return
	}
	fmt.Printf("The sum of two numbers is %d", add(number, number1))

}

func add(num int, num1 int) int {
	return num - -num1
}
