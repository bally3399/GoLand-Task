package main

import "fmt"

func main() {
	fmt.Println("N\tN2\tN3\tN4")
	for index := 1; index < 5; index++ {
		fmt.Print(index, "\t")
		fmt.Print(index*index, "\t")
		fmt.Print(index*index*index, "\t")
		fmt.Print(index * index * index * index)
		fmt.Println()
	}
}
