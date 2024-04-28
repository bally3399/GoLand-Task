package main

import "fmt"

func main() {
	var accountNumber string
	var balanceAtBeginning, charges, credit, limit, balance, newBalance int

	fmt.Print("Enter Account number: ")
	_, err2 := fmt.Scanln(&accountNumber)
	if err2 != nil {
		return
	}

	fmt.Print("Enter Balance at the beginning of the month: ")
	fmt.Scanln(&balanceAtBeginning)

	fmt.Print("Enter total of all items charged: ")
	_, err := fmt.Scanln(&charges)
	if err != nil {
		return
	}

	fmt.Print("Enter total of all credit applied: ")
	fmt.Scanln(&credit)

	fmt.Print("Enter allowed credit limit: ")
	fmt.Scanln(&limit)

	balance = balanceAtBeginning + charges
	newBalance = balance - credit

	if newBalance < limit {
		fmt.Println("Credit limit exceeded")
	} else {
		fmt.Println("New balance:", newBalance)
	}
}
