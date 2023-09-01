package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num uint8 = uint8(rand.Intn(3))

	switch num {
	case 0:
		fmt.Println("Is zero")
	case 1:
		fmt.Println("Is one")
		fallthrough
	case 3, 4:
		fmt.Println("3 or 4 were never meant to be")
	default:
		fmt.Println("Is two")
	}
}
