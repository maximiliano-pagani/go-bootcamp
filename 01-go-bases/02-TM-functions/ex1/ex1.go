package ex1

import "fmt"

func Ex1() {
	fmt.Println(getSalaryTaxes(40000))
	fmt.Println(getSalaryTaxes(60000))
	fmt.Println(getSalaryTaxes(160000))
	return
}

func getSalaryTaxes(salary int) (tax float64) {
	if salary > 50000 {
		tax += 0.17
	}
	if salary > 150000 {
		tax += 0.1
	}

	return float64(salary) * tax
}
