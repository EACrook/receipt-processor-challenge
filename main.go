package main

import (
	"github.com/gin-gonic/gin"
	"receipt-processor-challenge/api"
)



func main() {
	
	router := gin.Default()

	router.POST("/receipts/process", api.ProcessReceiptHandler)
	router.GET("/receipts/:id/points", api.GetPointsHandler)

	router.Run(":8080")
}
