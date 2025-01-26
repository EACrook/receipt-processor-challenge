package api

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"receipt-processor-challenge/models"
	"receipt-processor-challenge/store"
)

var s = store.NewStore()

func ProcessReceiptHandler(c *gin.Context) {
	var newReceipt models.StandardReceipt

	// Use the request processing function to bind and validate the receipt
	if err := BindAndValidateReceipt(c, &newReceipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid"})
		return
	}

	newReceipt = SetNewReceiptAndPointData(newReceipt)

	// Return the new receipt ID
	c.JSON(http.StatusOK, gin.H{"id": newReceipt.ID})
}

func GetPointsHandler(c *gin.Context) {
	id := c.Param("id")
	points, exists := s.GetPoints(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that ID"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"points": points.Points})
}