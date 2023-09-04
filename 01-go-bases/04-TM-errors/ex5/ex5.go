package ex5

import (
	"errors"
	"fmt"
)

// Vamos a hacer que nuestro programa sea un poco más complejo y útil.
// Desarrollá las funciones necesarias para permitir a la empresa calcular:
// Salario mensual de un trabajador según la cantidad de horas trabajadas.
// - La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
// - Dicha función deberá retornar más de un valor (salario calculado y error).
// - En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar
// el 10 % en concepto de impuesto.
// - En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo,
// la función debe devolver un error. El mismo tendrá que indicar “Error: el trabajador no puede haber
// trabajado menos de 80 hs mensuales”.

var (
	ErrNotEnoughHours = NotEnoughHoursError{}
	ErrInvalidHourFee = errors.New("Error: el valor por hora es inválido")
)

type NotEnoughHoursError struct{}

func (e *NotEnoughHoursError) Error() string {
	return fmt.Sprint("Error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
}

func calculateMonthlySalary(workedHours, feePerHour int) (float32, error) {
	if workedHours < 80 {
		return 0, &ErrNotEnoughHours
	}

	if feePerHour <= 0 {
		return 0, ErrInvalidHourFee
	}

	salary := float32(workedHours * feePerHour)
	if salary >= 150000 {
		salary *= 0.9
	}

	return salary, nil
}

func Ex5() {
	inputSampleData := [][]int {{100, 20}, {100, 2000}, {70, 35}, {120, -5}}
	
	var (
		salary float32
		err error
	)

	for _, data := range inputSampleData {
		salary, err = calculateMonthlySalary(data[0], data[1])

		if err != nil {
			fmt.Println(data[0], data[1], err)
		} else {
			fmt.Println(data[0], data[1], "El salario es de", salary)
		}
	}
}
