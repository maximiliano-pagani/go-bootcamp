package ex2

import (
	"errors"
	"fmt"
)

// En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
// Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: el salario es menor a 10.000"
// y lanzalo en caso de que “salary” sea menor o igual a  10.000. La validación debe ser hecha con la función “Is()” dentro del “main”.

var ErrLowSalary LowSalaryError = LowSalaryError{}

type LowSalaryError struct{}

func (e *LowSalaryError) Error() string {
	return fmt.Sprint("Error: el salario es menor a 10.000")
}

func salaryPaysTax(salary int) (bool, error) {
	if salary <= 10000 {
		return false, &ErrLowSalary
	}
	return true, nil
}

func Ex2() {
	salariesSampleData := []int{8000, 50000}

	var err error

	for _, salary := range salariesSampleData {
		_, err = salaryPaysTax(salary)
		
		if errors.Is(err, &ErrLowSalary) {
			fmt.Println(salary, err)
		} else {
			fmt.Println(salary, "Debe pagar impuesto")
		}
	}
}
