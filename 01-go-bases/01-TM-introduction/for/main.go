package main

import (
	"fmt"
)

func main() {
	for i := 0; i != 5; i++ {
		fmt.Print(i)
	}

	var i int = 4

	for i != 9 {
		i++
		fmt.Print(i)
	}

	var fruits = []string{"apple", "banana", "orange"}

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}
