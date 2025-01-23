package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DateOnly struct {
	time.Time
}

type TimeOnly struct {
	time.Time
}

type Store struct {
	// ensures proper synchronization for concurrent processing -- ensures only one goroutine can access the resource at a time
	mu sync.Mutex
	receipts map[string]StandardReceipt
}

type StandardReceipt struct {
	ID string `json:"id"`
	Retailer     string    `json:"retailer"`
}

func NewStore() *Store {
	return &Store{
		receipts: make(map[string]StandardReceipt),
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

func main() {
	store := NewStore()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the API!"})
	})

	router.POST("/receipts/process", func(c *gin.Context) {
		var newReceipt StandardReceipt

		if err := c.ShouldBindJSON(&newReceipt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newReceipt.ID = uuid.New().String()

		store.AddReceipt(newReceipt)

		c.JSON(http.StatusCreated, newReceipt)
	})

	router.GET("/receipts/:id", func(c *gin.Context) {
		id := c.Param("id")
		receipt, exists := store.GetReceipt(id)
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "receipt not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"points": receipt})
	})

	router.Run(":8080")
}
