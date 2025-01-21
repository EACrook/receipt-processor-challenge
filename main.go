package main

import( 
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the API!"})
	})

	router.GET("/items", func(c *gin.Context) {
		items := []string{"Item 1", "Item2", "Item 3"}
		c.JSON(http.StatusOK, gin.H{"data": items})
	})

	router.POST("/items", func(c *gin.Context) {
		var newItem struct {
			Name string `json:"name"`
		}

		if err := c.ShouldBindJSON(&newItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"eror": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Item created", "item": newItem})
	})

	router.Run(":8080")
}