package calculator

import (
	"errors"
	"math"
)

func Add(inputs ...float64) float64 {
	var result float64
	for _, input := range inputs {
		result += input
	}
	return result
}

func Subtract(inputs ...float64) float64 {
	result := 2 * inputs[0]
	for _, input := range inputs {
		result -= input
	}
	return result
}

func Multiply(inputs ...float64) float64 {
	var result float64 = 1
	for _, input := range inputs {
		result *= input
	}
	return result
}

func Divide(inputs ...float64) (float64, error) {
	result := inputs[0]
	for i, input := range inputs {
		if i == 0 {
			continue
		}
		if input == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= input
	}
	return result, nil

	//if b == 0 {
	//}
	//return a / b, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of negative number")
	}
	return math.Sqrt(a), nil
}
