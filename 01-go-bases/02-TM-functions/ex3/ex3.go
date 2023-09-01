package ex3

import "fmt"

const (
	catA = "A"
	catB = "B"
	catC = "C"
)

func Ex3() {
	fmt.Println("Init ex3")
	fmt.Println(getCategorySalary(60, catA))
	fmt.Println(getCategorySalary(60, catB))
	fmt.Println(getCategorySalary(60, catC))

	return
}

func getCategorySalary(minutes int, cat string) float32 {
	var salary float32
	switch cat {
	case catA:
		salary = calculateSalary(minutes, 3000, 0.5)
	case catB:
		salary = calculateSalary(minutes, 1500, 0.2)
	case catC:
		salary = calculateSalary(minutes, 1000, 0.0)
	}

	return salary
}

func calculateSalary(minutes int, moneyPerHour int, bonus float32) float32 {
	var baseSalary float32 = (float32(minutes) / 60.0) * float32(moneyPerHour)
	return baseSalary + baseSalary*bonus
}
