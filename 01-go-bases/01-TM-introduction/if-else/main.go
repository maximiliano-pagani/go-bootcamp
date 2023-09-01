package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num uint8 = uint8(rand.Intn(3))

	if x := uint8(0); !(num != x) {
		fmt.Println("Is zero")
	} else if num != 2 {
		fmt.Println("Is one")
	} else {
		fmt.Println("Is two")
	}
}
