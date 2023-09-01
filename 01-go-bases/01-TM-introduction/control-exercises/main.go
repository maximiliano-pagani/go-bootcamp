package main

import (
	"fmt"
)

func main() {
	var phrase string = "Hola mundo"

	fmt.Println("Length is", len(phrase))
	for _, char := range phrase {
		fmt.Print(string(char))
	}

	fmt.Println()
	fmt.Println()

	isElegibleForLoan(21, false, 1, 40000)
	isElegibleForLoan(23, false, 1, 40000)
	isElegibleForLoan(23, true, 1, 40000)
	isElegibleForLoan(23, true, 2, 400000)
	isElegibleForLoan(23, true, 2, 10000)

	fmt.Println()
	fmt.Println()

	var employees = map[string]int{"Benja": 20, "Brenda": 26, "Dario": 44, "Pedro": 30}

	fmt.Println("Benja tiene", employees["Benja"], "años")

	var count int = 0
	for _, age := range employees {
		if age > 21 {
			count++
		}
	}

	fmt.Println("Hay", count, "empleados mayores de 21 años")
	delete(employees, "Pedro")
}

func isElegibleForLoan(age int, isEmployee bool, workingTimeInYears int, salary int) {
	switch {
	case age <= 22:
		fmt.Println("Menor de 22 años")
	case !isEmployee:
		fmt.Println("No empleado")
	case workingTimeInYears <= 1:
		fmt.Println("Antiguedad menor a un año")
	case salary > 100000:
		fmt.Println("Debe pagar intereses")
	default:
		fmt.Println("Puede recibir el préstamo")
	}
}
