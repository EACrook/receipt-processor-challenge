package helpers

import (
	"unicode"
	"strconv"
	"time"
	"strings"
	"math"

	"receipt-processor-challenge/models"
)	

func PointCalculationAll(receipt models.StandardReceipt) int {
	retailerPoints := pointCalculationRetailer(receipt.Retailer)
	receiptTotalPoints := pointsCalculationReceiptTotal(receipt.Total)
	lineItemsPoints := pointsCalculationLineItems(receipt.Items)
	datePoints := pointCalculationDate(receipt.PurchaseDate)
	timePoints := pointCalculationTime(receipt.PurchaseTime)

	return retailerPoints + receiptTotalPoints + lineItemsPoints + datePoints + timePoints
}

func pointCalculationRetailer(retailer string) int {
	points := 0
	for _, char := range retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points++
		}
	}
	return points
}

func pointsCalculationReceiptTotal(total string) int {
	points := 0
	i, err := strconv.ParseFloat(total, 64)
	if err != nil {
		panic(err)
	}

	if isRoundNumber(i) {
		points += 50 
	}
	if isMultipleFloat(i, 0.25) {
		points += 25
	}
	return points
}

func pointsCalculationLineItems(items []models.ItemData) int {
	points := 0
	points += groupedByTwoPoints(len(items))
	points += itemDescriptionPoints(items)
	return points
}

func pointCalculationDate(dateStr string) int {
	points := 0
	layout := "2006-01-02"
	
	date, err := time.Parse(layout, dateStr) 
	if err != nil {
		panic(err)
	}

	day := date.Day()

	if !isMultipleInt(day, 2) {
		return points + 6
	}
	return points	
}

func pointCalculationTime(timeStr string) int {
	points := 0
	layout := "15:04"

	purchaseTime, err := time.Parse(layout, timeStr)
	if err != nil {
		panic(err)
	}
	startTime, _ := time.Parse(layout, "14:00")
	endTime, _ := time.Parse(layout, "16:00") 

	if purchaseTime.After(startTime) && purchaseTime.Before(endTime) {
		return points + 10
	}
	return points
}

func groupedByTwoPoints(itemsLength int) int {
	if !isMultipleInt(itemsLength, 2) {
		itemsLength = itemsLength - 1
	}
	return itemsLength/2 * 5
}

func itemDescriptionPoints(items []models.ItemData) int {
	points := 0
	for _, item := range items {
		trimmed := strings.TrimSpace(item.ShortDescription)
		if isMultipleInt(len(trimmed), 3) {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				panic(err)
			}
			points += int(math.Ceil(price * 0.2))
		}
	}
	return points
}