package main

import "fmt"

func main() {
	for index := 1; index <= 10; index++ {
		for input := 1; input <= index; input++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()
	for index := 1; index <= 10; index++ {
		for input := index; input <= 10; input++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()
	for index := 1; index <= 10; index++ {
		for input := index; input <= 10; input++ {
			fmt.Print(" ")
		}
		for input1 := 1; input1 <= index; input1++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()
	for input := 1; input <= 10; input++ {
		for index := 1; index <= 10; index++ {
			fmt.Print(" ")
		}
		for input1 := input; input1 <= 10; input1++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
