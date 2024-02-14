package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	value1 := "10"
	value2 := "5.5"
	operation := "+"
	result, _ := calculate(value1, value2, operation)
	fmt.Println(result)
}

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) (float64, error) {
	input1Converted := convertInputToValue(input1)
	input2Converted := convertInputToValue(input2)
	switch operation {
	case "+":
		return addValues(input1Converted, input2Converted), nil
	case "-":
		return subtractValues(input1Converted, input2Converted), nil
	case "*":
		return multiplyValues(input1Converted, input2Converted), nil
	case "/":
		if input2Converted == 0 {
			return 0, errors.New("division by zero")
		}
		return divideValues(input1Converted, input2Converted), nil
	default:
		return 0, fmt.Errorf("invalid operation: %v", operation)
	}
}

func convertInputToValue(input string) float64 {
	f, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a number", input)
		panic(message)
	}
	return f
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
