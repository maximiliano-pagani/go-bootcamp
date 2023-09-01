package ex4

import (
	"fmt"
	"slices"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func Ex4() {
	fmt.Println("Init ex4")
	minFunc, minErr := operation(minimum)
	averageFunc, avgErr := operation(average)
	maxFunc, maxErr := operation(maximum)
	_, invErr := operation("invalidOperation")

	fmt.Println(minFunc(2, 3, 3, 4, 10, 2, 4, 5), minErr)
	fmt.Println(averageFunc(2, 3, 3, 4, 1, 2, 4, 5), avgErr)
	fmt.Println(maxFunc(2, 3, 3, 4, 1, 2, 4, 5), maxErr)
	fmt.Println(invErr)

	return
}

func operation(opCode string) (func(values ...int) float64, string) {
	switch opCode {
	case minimum:
		return minOperation, ""
	case average:
		return avgOperation, ""
	case maximum:
		return maxOperation, ""
	default:
		return nil, "Cálculo no definido"
	}
}

func minOperation(values ...int) float64 {
	return float64(slices.Min(values))
}

func avgOperation(values ...int) float64 {
	var count float64
	for _, value := range values {
		count += float64(value)

	}

	return count / float64(len(values))
}

func maxOperation(values ...int) float64 {
	return float64(slices.Max(values))
}
