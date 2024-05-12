package main

import "fmt"

func main() {
	for index := 1; index <= 10; index++ {
		for idx := 1; idx <= 10-index; idx++ {
			fmt.Print(" ")
		}
		for idx := 1; idx <= 2*index-1; idx++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for index := 10 - 1; index >= 1; index-- {
		for idx := 1; idx <= 10-index; idx++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*index-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
