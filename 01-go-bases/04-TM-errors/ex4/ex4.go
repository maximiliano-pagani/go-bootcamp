package ex4

import (
	"errors"
	"fmt"
)

// Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por
// parámetro el valor de “salary” indicando que no alcanza el mínimo imponible (el mensaje mostrado por consola
// deberá decir: “Error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”, siendo
// [salary] el valor de tipo int pasado por parámetro).

var ErrLowSalary error = errors.New("Error: el mínimo imponible es de 150.000")

func salaryPaysTax(salary int) (bool, error) {
	if salary <= 150000 {
		return false, fmt.Errorf("%w y el salario ingresado es de: %d", ErrLowSalary, salary)
	}
	return true, nil
}

func Ex4() {
	salariesSampleData := []int{50000, 200000}

	var err error

	for _, salary := range salariesSampleData {
		_, err = salaryPaysTax(salary)

		if errors.Is(err, ErrLowSalary) {
			fmt.Println(salary, err)
		} else {
			fmt.Println(salary, "Debe pagar impuesto")
		}
	}
}
