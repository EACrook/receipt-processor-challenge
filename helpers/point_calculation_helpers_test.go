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
				t.Errorf("For input %s, expected %d but got %d", test.input, test.expected, result)
			}
		})
	}
}

func TestPointsCalculationReceiptTotal(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected int
	} {
		{"Test total is round and is multiple of 0.25", "50.00", 75},
		{"Test total is only a multiple of 0.25", "1.25", 25},
		{"Test 0.00 will return 75 since 0 is a round number and a multiple of 0.25", "0.00", 75},
		{"Test a total that will return 0 points", "19.99", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := pointsCalculationReceiptTotal(test.input)
			if result != test.expected {
				t.Errorf("For total %s, expected %d but got %d", test.input, test.expected, result)
			}
		})
	}
}

func TestPointCalculationDate(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected int
	} {
		{"An odd day will return 6 points", "1996-01-05", 6},
		{"An even day will return 0 points", "1996-01-06", 0},
		{"Invalid date will return 0", "not-a-date", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := pointCalculationDate(test.input)
			if result != test.expected {
				t.Errorf("For date %s, expected %d but got %d", test.input, test.expected, result)
			}
		})
	}
}

func TestPointCalculationTime(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected int
	} {
		{"Time of purchase is after 2PM and before 4PM", "15:59", 10},
		{"Time of purchase is outside of 2PM and 4PM", "8:48", 0},
		{"Invalid time will return 0", "not-a-time", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := pointCalculationTime(test.input)
			if result != test.expected {
				t.Errorf("For date %s, expected %d but got %d", test.input, test.expected, result)
			}
		})
	}
}

func TestGroupedByTwoPoints(t *testing.T) {
	tests := []struct {
		name string
		input int
		expected int
	} {
		{"Test an odd number length of items", 29, 70},
		{"Test an even number length of items", 6, 15},
		{"Test 0 will return 0", 0, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := groupedByTwoPoints(test.input)
			if result != test.expected {
				t.Errorf("For item length %d, expected %d but got %d", test.input, test.expected, result)
			}
		})
	}
}

var testItemData1 = []models.ItemData{
	{
		ShortDescription: "Mountain Dew 12PK",
		Price:            "6.49",
	},
	{
		ShortDescription: "Emils Cheese Pizza",
		Price:            "12.25",
	},
	{
		ShortDescription: "Knorr Creamy Chicken",
		Price:            "1.26",
	},
	{
		ShortDescription: "Doritos Nacho Cheese",
		Price:            "3.35",
	},
	{
		ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
		Price:            "12.00",
	},
}

var testItemData2 = []models.ItemData{
	{
		ShortDescription: "Gatorade",
		Price: "2.25",
	},
	{
		ShortDescription: "Gatorade",
		Price: "2.25",
	},
	{
		ShortDescription: "Gatorade",
		Price: "2.25",
	},
	{
		ShortDescription: "Gatorade",
		Price: "2.25",
	},
}

var testItemDataEmpty = []models.ItemData{
	{},
}

func TestItemDescriptionPoints(t *testing.T) {
	tests := []struct {
		name string
		input []models.ItemData
		expected int
	} {
		{"Test item data 1", testItemData1, 6},
		{"Test item data 2", testItemData2, 0},
		{"Test empty item data", testItemDataEmpty, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := itemDescriptionPoints(test.input)
			if result != test.expected {
				t.Errorf("For item %v, expected %d but got %d", test.input, test.expected, result)
			}
		})
	}
}

func TestPointsCalculationLineItems(t *testing.T) {
	tests := []struct {
		name string
		input []models.ItemData
		expected int
	} {
		{"Test item data 1", testItemData1, 16},
		{"Test item data 2", testItemData2, 10},
		{"Test empty item data", testItemDataEmpty, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := pointsCalculationLineItems(test.input)
			if result != test.expected {
				t.Errorf("For item %v, expected %d but got %d", test.input, test.expected, result)
			}
		})
	}
}

var testReceiptData1 = models.StandardReceipt{
	ID: "123",
	Retailer: "Target",
	PurchaseDate: "2022-01-01",
	PurchaseTime: "13:01",
	Items: []models.ItemData{
		{
			ShortDescription: "Mountain Dew 12PK",
			Price: "6.49",
		},
		{
			ShortDescription: "Emils Cheese Pizza",
			Price: "12.25",
		},
		{
			ShortDescription: "Knorr Creamy Chicken",
			Price: "1.26",
		},
		{
			ShortDescription: "Doritos Nacho Cheese",
			Price: "3.35",
		},
		{
			ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
			Price: "12.00",
		},
	},
	Total: "35.35",
}

var testReceiptData2 = models.StandardReceipt{
	ID: "456",
	Retailer: "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "14:33",
	Items: []models.ItemData{
		{
			ShortDescription: "Gatorade",
			Price: "2.25",
		},
		{
			ShortDescription: "Gatorade",
			Price: "2.25",
		},
		{
			ShortDescription: "Gatorade",
			Price: "2.25",
		},
		{
			ShortDescription: "Gatorade",
			Price: "2.25",
		},
	},
	Total: "9.00",
}

var testReceiptData3 = models.StandardReceipt{
	ID: "789",
	Retailer: "Walgreens",
	PurchaseDate: "2022-01-02",
	PurchaseTime: "08:13",
	Items: []models.ItemData{
		{
			ShortDescription: "Pepsi - 12-oz",
			Price: "1.25",
		},
		{
			ShortDescription: "Dasani", 
			Price: "1.40",
		},
	},
	Total: "2.65",
}

var testReceiptData4 = models.StandardReceipt{
	ID: "741",
	Retailer: "Target",
	PurchaseDate: "2022-01-02",
	PurchaseTime: "13:13",
	Items: []models.ItemData{
		{
			ShortDescription: "Pepsi - 12-oz", 
			Price: "1.25",
		},
	},
	Total: "1.25",
}

var testEmptyReceipt = models.StandardReceipt{
	ID: "852",
	Retailer: "",
	PurchaseDate: "",
	PurchaseTime: "",
	Items: []models.ItemData{
		{
			ShortDescription: "", 
			Price: "",
		},
	},
	Total: "",
}

func TestPointCalculationAll(t *testing.T) {
	tests := []struct {
		name string
		input models.StandardReceipt
		expected int
	} {
		{"Test receipt 1", testReceiptData1, 28},
		{"Test receipt 2", testReceiptData2, 109},
		{"Test receipt 3", testReceiptData3, 15},
		{"Test receipt 4", testReceiptData4, 31},
		{"Test empty receipt", testEmptyReceipt, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PointCalculationAll(test.input)
			if result != test.expected {
				t.Errorf("For receipt %v, expected %d but got %d", test.input, test.expected, result)
			}
		})
	}
}