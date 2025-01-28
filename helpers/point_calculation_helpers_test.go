package helpers

import (
	"receipt-processor-challenge/models"
	"testing"
)

func TestPointCalculationRetailer(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected int
	} {
		{"Test only characters name", "Walgreens", 9},
		{"Test with characters and symbols", "M&M Corner Market", 14},
		{"Test with characters and numbers", "PunchOut2Go", 11},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := pointCalculationRetailer(test.input)
			if result != test.expected {
				t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
			}
		})
	}
}

func TestPointsCalculationLineItems(t *testing.T) {
	test := []struct {
		name string
		input []models.ItemData
		expected int
	} {
		
	}
}
