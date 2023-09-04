package ex1

import "fmt"

// En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
// Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: el salario ingresado no alcanza el
// mínimo imponible" y lanzalo en caso de que “salary” sea menor a 150.000. De lo contrario, tendrás que imprimir por consola el
// mensaje “Debe pagar impuesto”.

type LowSalaryError struct{}

func (e *LowSalaryError) Error() string {
	return fmt.Sprint("Error: el salario ingresado no alcanza el mínimo imponible")
}

func salaryPaysTax(salary int) (bool, error) {
	if salary < 150000 {
		return false, &LowSalaryError{}
	}
	return true, nil
}

func Ex1() {
	salariesSampleData := []int{50000, 200000}

	var err error

	for _, salary := range salariesSampleData {
		_, err = salaryPaysTax(salary)
		
		if err != nil {
			fmt.Println(salary, err)
		} else {
			fmt.Println(salary, "Debe pagar impuesto")
		}
	}
}
