package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	value1 := "10"
	value2 := "5.5"
	operation := "+"
	result := calculate(value1, value2, operation)
	fmt.Println(result)
}

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
	input1Converted := convertInputToValue(input1)
	input2Converted := convertInputToValue(input2)
	result := 0.0
	if operation == "+" {
		result = addValues(input1Converted, input2Converted)
		return result
	} else if operation == "-" {
		result = subtractValues(input1Converted, input2Converted)
		return result
	} else if operation == "*" {
		result = multiplyValues(input1Converted, input2Converted)
		return result
	} else if operation == "/" {
		result = divideValues(input1Converted, input2Converted)
		return result
	} else {
		panic("Invalid operation")
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
