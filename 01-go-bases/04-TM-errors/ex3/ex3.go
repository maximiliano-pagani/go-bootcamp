package ex3

import (
	"errors"
	"fmt"
)

// Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que,
// en reemplazo de “Error()”,  se implemente “errors.New()”.

var ErrLowSalary error = errors.New("Error: el salario es menor a 10.000")

func salaryPaysTax(salary int) (bool, error) {
	if salary <= 10000 {
		return false, ErrLowSalary
	}
	return true, nil
}

func Ex3() {
	salariesSampleData := []int{8000, 50000}

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