package main

import "fmt"

func main() {
	var taxRate float64
	counter := 0
	total := 0.0

	fmt.Print("Enter tax rate(use -1 to exit): ")
	fmt.Scanln(&taxRate)

	for taxRate != -1 {
		total = taxRate * 15 / 100
		counter = counter + 1
		fmt.Print("Enter tax rate(use -1 to exit): ")
		fmt.Scanln(&taxRate)

	}
	totalTaxRate := total * 0.2
	fmt.Println(totalTaxRate)

}
