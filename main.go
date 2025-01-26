package main

import (
	"net/http"
	"regexp"
	"sync"
	"time"
	"unicode"
	"strconv"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Store struct {
	// ensures proper synchronization for concurrent processing -- ensures only one goroutine can access the resource at a time
	mu sync.Mutex
	receipts map[string]StandardReceipt
	points map[string]Points
}

type StandardReceipt struct {
	ID string `json:"id"`
	Retailer string `json:"retailer" binding:"required,retailerFormat"`
	PurchaseDate string `json:"purchaseDate" binding:"required,dateFormat"`
	PurchaseTime string `json:"purchaseTime" binding:"required,timeFormat"`
	Items []ItemData `json:"items" binding:"required,dive"`
	Total string `json:"total" binding:"required,priceFormat"`
}

type ItemData struct {
	ShortDescription string `json:"shortDescription" binding:"required"`
	Price string `json:"price" binding:"required,priceFormat"`
}

type Points struct {
	ID string `json:"id"`
	Points int `json:"points"`
}

func NewStore() *Store {
	return &Store{
		receipts: make(map[string]StandardReceipt),
		points: make(map[string]Points),
	}
}

func (s *Store) AddReceipt(receipt StandardReceipt) {
	// prevents other goroutines from accessing the resource
	s.mu.Lock()
	// allows other goroutines to access the resource
	defer s.mu.Unlock()
	s.receipts[receipt.ID] = receipt
}

func (s *Store) GetReceipt(id string) (StandardReceipt, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	receipt, exists := s.receipts[id]
	return receipt, exists
}

func (s *Store) AddPoints(points Points) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.points[points.ID] = points
}

func (s *Store) GetPoints(id string) (Points, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	points, exists := s.points[id]
	return points, exists
}

// Validates that our price and total fields match the correct format
var validatePrice validator.Func = func (fl validator.FieldLevel) bool {
	pattern := "^\\d+\\.\\d{2}$"

	matched, _ := regexp.MatchString(pattern, fl.Field().String())
	return matched
}

var validateRetailer validator.Func = func (fl validator.FieldLevel) bool {
	pattern := "^[\\w\\s\\-&]+$"

	matched, _ := regexp.MatchString(pattern, fl.Field().String())
	return matched
}

var validateDate validator.Func = func (fl validator.FieldLevel) bool {
	const layout = "2006-01-02" // Date format: YYYY-mm-dd

	// TODO: I don't like this -- try to see if I can match similar formats to the rest
	_, err := time.Parse(layout, fl.Field().String())
	return err == nil
}

var validateTime validator.Func = func (fl validator.FieldLevel) bool {
	const layout = "15:04" // Time format: HH:mm

	_, err := time.Parse(layout, fl.Field().String())
	return err == nil
}

var validateDescription validator.Func = func (fl validator.FieldLevel) bool {
	pattern := "^[\\w\\s\\-]+$"

	matched, _ := regexp.MatchString(pattern, fl.Field().String())
	return matched
}

func pointCalculationAll(receipt StandardReceipt) int {
	retailerPoints := pointCalculationRetailer(receipt.Retailer)
	receiptTotalPoints := pointsCalculationReceiptTotal(receipt.Total)
	lineItemsPoints := pointsCalculationLineItems(receipt.Items)

	return retailerPoints + receiptTotalPoints + lineItemsPoints
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

func pointsCalculationLineItems(items []ItemData) int {
	points := 0
	points += groupedByTwoPoints(len(items))
	points += itemDescriptionPoints(items)
	return points
}

func isRoundNumber(total float64) bool {
	return total == math.Floor(total)
}

func isMultipleFloat(dividend float64, divisor float64) bool {
	remainder := math.Mod(dividend, divisor)
	return math.Abs(remainder) < 1e-9
}

func isMultipleInt(dividend int, divisor int) bool {
	return dividend%divisor == 0
}

func groupedByTwoPoints(itemsLength int) int {
	if !isMultipleInt(itemsLength, 2) {
		itemsLength = itemsLength - 1
	}
	return itemsLength/2 * 5
}

func itemDescriptionPoints(items []ItemData) int {
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

func main() {
	store := NewStore()
	router := gin.Default()

	// validator for price and total
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("priceFormat", validatePrice)
		v.RegisterValidation("retailerFormat", validateRetailer)
		v.RegisterValidation("dateFormat", validateDate)
		v.RegisterValidation("timeFormat", validateTime)
		v.RegisterValidation("descriptionFormat", validateDescription)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the API!"})
	})

	router.POST("/receipts/process", func(c *gin.Context) {
		var newReceipt StandardReceipt
		var points Points

		if err := c.ShouldBindJSON(&newReceipt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newId := uuid.New().String()
		newReceipt.ID = newId
		points.ID = newId
		points.Points = pointCalculationAll(newReceipt) 

		store.AddReceipt(newReceipt)
		store.AddPoints(points)
		c.JSON(http.StatusCreated, newReceipt)
	})

	router.GET("/receipts/:id", func(c *gin.Context) {
		id := c.Param("id")
		receipt, exists := store.GetReceipt(id)
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "receipt not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"receipt": receipt})
	})

	// TODO: write functions for points maths
	router.GET("receipts/:id/points", func(c *gin.Context) {
		id := c.Param("id")
		points, exists := store.GetPoints(id)
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "points not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"points": points})
	})

	router.Run(":8080")
}
