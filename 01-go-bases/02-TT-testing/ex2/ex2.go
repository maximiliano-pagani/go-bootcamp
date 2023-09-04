package ex2

import "fmt"

func Ex2() {
	fmt.Println(getAverage(5.5, 2.0, 2.0))
	fmt.Println(getAverage(5.5))
	fmt.Println(getAverage(7.5, -2.0))
	fmt.Println(getAverage())
}

func getAverage(grades ...float32) float32 {
	totalGrades := len(grades)

	if totalGrades <= 0 {
		return 0
	}

	var count float32
	for _, grade := range grades {
		if grade > 0 {
			count += grade
		} else {
			totalGrades--
		}

	}

	return count / float32(totalGrades)
}
