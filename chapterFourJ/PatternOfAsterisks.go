package main

import "fmt"

func main() {
	for index := 1; index <= 16; index++ {
		for input := 1; input <= 16; input++ {
			if (index+input)%2 == 0 {
				fmt.Print("* ")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
