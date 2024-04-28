package main

import (
	"fmt"
)

func main() {
	var digit int
	fmt.Println("Enter five-digit number to check for palindrome:")
	fmt.Scanln(&digit)

	if digit < 10000 || digit > 99999 {
		fmt.Println("Error message!! Please enter a five-digit number.")
		return
	}

	result := isPalindrome(digit)
	fmt.Println("Is it a palindrome?", result)
}

func isPalindrome(number int) bool {
	originalNumber := number
	reverse := 0
	for number > 0 {
		lastDigit := number % 10
		reverse = reverse*10 + lastDigit
		number = number / 10
	}

	return originalNumber == reverse
}
