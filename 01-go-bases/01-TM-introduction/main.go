package main

import "fmt"

func main() {
	fmt.Println("Init 01-introduction module")

	// Hello world
	fmt.Println("Hello world")
	var username = "admin"
	fmt.Println("Hello,", username)
	fmt.Println()

	// Exercise 1
	fmt.Println("Init ex1")
	var name, lastName string = "Maximiliano", "Pagani"
	fmt.Println("Hello,", name, lastName)
	fmt.Println()

	// Exercise 2
	fmt.Println("Init ex2")
	var (
		temp     float64 = 19.0
		humidity float64 = 0.45
		pressure float64 = 1023.6
	)

	fmt.Printf("%f C, %f %%, %f HPA\n", temp, humidity, pressure)
}
