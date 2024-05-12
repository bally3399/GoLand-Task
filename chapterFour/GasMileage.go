package main

import "fmt"

func main() {
	var trip int
	totalMilesDriven := 0.0
	totalGallonsUsed := 0.0

	for trip != -1 {
		fmt.Print("Enter numbers of miles: ")
		var miles float64
		_, err := fmt.Scanln(&miles)
		if err != nil {
			return
		}

		fmt.Print("Enter numbers of gallons: ")
		var gallons float64
		_, err = fmt.Scanln(&gallons)
		if err != nil {
			return
		}

		totalMilesDriven += miles
		totalGallonsUsed += gallons

		milesPerGallon := miles / gallons
		fmt.Println(milesPerGallon)

		fmt.Print("Enter numbers of trip (Enter -1 to exit): ")
		_, err = fmt.Scanln(&trip)
		if err != nil {
			return
		}
	}

	totalMilesPerGallon := totalMilesDriven / totalGallonsUsed
	fmt.Printf("Total miles per gallon is: %.2f\n", totalMilesPerGallon)
}
