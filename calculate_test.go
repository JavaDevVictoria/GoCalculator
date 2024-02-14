package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const value1 = "10"
const value2 = "5.5"
const operation = "+"

func TestCalculate(t *testing.T) {
	assert.Equal(t, calculate(value1, value2, operation), 15.5, "Result is incorrect")
}

func TestCalculateTable(t *testing.T) {
	tests := []struct {
		name      string
		input1    string
		input2    string
		operation string
		want      float64
	}{
		{"Addition", "10", "5.5", "+", 15.5},
		{"Subtraction", "10", "5.5", "-", 4.5},
		{"Multiplication", "10", "5.5", "*", 55},
		{"Division", "10", "5.5", "/", 1.8181818181818181},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.input1, tt.input2, tt.operation); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertInputToValue(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  float64
	}{
		{"Positive number", "10", 10},
		{"Negative number", "-5.5", -5.5},
		{"Zero", "0", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertInputToValue(tt.input); got != tt.want {
				t.Errorf("convertInputToValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
