package main

import "fmt"

func main() {
	total := 0.0
	count := 0
	var item float64
	for item != -1 {
		fmt.Print("Enter number of items (Enter -1 to exit): ")
		fmt.Scanln(&item)
		if item != -1 {
			total += item
			count++
		}
	}

	amountEarned := total * 0.09
	amount := amountEarned + 200
	fmt.Println(amount)
}
