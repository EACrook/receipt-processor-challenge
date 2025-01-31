package store

import (
	"testing"

	"receipt-processor-challenge/models"

)

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

var testPoints1 = models.Points{
	ID: "123",
	Points: 28,
}

func TestStore(t *testing.T) {
	store := NewStore()

	store.AddReceipt(testReceiptData1)
	store.AddPoints(testPoints1)

	storedPoints, _ := store.GetPoints("123")

	if storedPoints.ID != "123" {
		t.Errorf("expected ID '123', got %v", storedPoints.ID)
	}

	if storedPoints.Points != 28 {
		t.Errorf("expected points to be 28 ', got %v", storedPoints.ID)
	}
}