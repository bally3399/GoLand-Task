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
	_, err := fmt.Scanln(&balanceAtBeginning)
	if err != nil {
		return
	}

	fmt.Print("Enter total of all items charged: ")
	_, err = fmt.Scanln(&charges)
	if err != nil {
		return
	}

	fmt.Print("Enter total of all credit applied: ")
	_, err2 = fmt.Scanln(&credit)
	if err2 != nil {
		return
	}

	fmt.Print("Enter allowed credit limit: ")
	_, err2 = fmt.Scanln(&limit)
	if err2 != nil {
		return
	}

	balance = balanceAtBeginning + charges
	newBalance = balance - credit

	if newBalance < limit {
		fmt.Println("Credit limit exceeded")
	} else {
		fmt.Println("New balance:", newBalance)
	}
}
