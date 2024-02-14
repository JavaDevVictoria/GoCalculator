package main

import (
	"testing"
)

func TestCalculateTable(t *testing.T) {
	tests := []struct {
		name      string
		input1    string
		input2    string
		operation string
		want      float64
		wantErr   bool
	}{
		{"Addition", "10", "5.5", "+", 15.5, false},
		{"Subtraction", "10", "5.5", "-", 4.5, false},
		{"Multiplication", "10", "5.5", "*", 55, false},
		{"Division", "10", "5.5", "/", 1.8181818181818181, false},
		{"Invalid operation", "10", "5.5", "%", 0, true},
		{"Divide by zero", "10", "0", "/", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculate(tt.input1, tt.input2, tt.operation)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
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
