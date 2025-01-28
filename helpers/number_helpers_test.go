package helpers

import (
	"testing"
)

func TestIsRoundNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected bool
	}{
		{"Positive whole number", 10.00, true},
		{"Negative whole number", -23.00, true},
		{"Zero", 0.00, true},
		{"Positive number with fractional part", 1.23, false},
		{"Negative number with fractional part", -3.33, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isRoundNumber(test.input)
			if result != test.expected {
				t.Errorf("For input %f, expected %v but got %v", test.input, test.expected, result)
			}
		})
	}
}

func TestIsMultipleFloat(t *testing.T) {
	tests := []struct {
		name          string
		inputDividend float64
		inputDivisor  float64
		expected      bool
	}{
		{"Test is a multiple", 12.00, .25, true},
		{"Test is not a multiple", 43.69, .50, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isMultipleFloat(test.inputDividend, test.inputDivisor)
			if result != test.expected {
				t.Errorf("For dividend %f and divisor %f, expected %v, but got %v", test.inputDividend, test.inputDivisor, test.expected, result)
			}
		})
	}
}

func TestIsMultipleInt(t *testing.T) {
	tests := []struct {
		name          string
		inputDividend int
		inputDivisor  int
		expected      bool
	}{
		{"Test is a multiple", 12, 3, true},
		{"Test is not a multiple", 14, -5, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isMultipleInt(test.inputDividend, test.inputDivisor)
			if result != test.expected {
				t.Errorf("For dividend %d and divisor %d, expected %v, but got %v", test.inputDividend, test.inputDivisor, test.expected, result)
			}
		})
	}
}
