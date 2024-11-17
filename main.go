package main

import (
	"net/http"
	"receipt-processor/lib"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/validator.v2"
)

const SERVER_ADDRESS = "localhost:8080"

var m map[string]int

func main() {
	initializeStorage()
	router := gin.Default()

	router.POST("/receipts/process", postReceipt)
	router.GET("/receipts/:id/points", getPoints)

	router.Run(SERVER_ADDRESS)
}

// Initialize storage
// placeholder for if / when a real storage method is needed
func initializeStorage() {
	m = make(map[string]int)
}

// Gets the point value of a scored receipt based on its id.
// If the id is not found returns an error
func getPoints(c *gin.Context) {
	id := c.Param("id")
	points, exists := m[id]
	if exists {
		c.JSON(http.StatusOK, gin.H{"points": points})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
	}
}

// Score a receipt and store the value with a unique id
// Returns the unique id
func postReceipt(c *gin.Context) {
	var newReceipt lib.Receipt

	// get the receipt from the request body
	if err := c.ShouldBindJSON(&newReceipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate the receipt fields
	if err := validator.Validate(newReceipt); err != nil {
		error := strings.Replace(err.Error(), "regular expression mismatch", "invalid value", 1)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": error})
		return
	}

	// score the receipt
	score := lib.ScoreReceipt(&newReceipt)

	// generate an id
	id := uuid.NewString()

	// store the id with the points
	m[id] = score

	// return the id
	c.JSON(http.StatusOK, gin.H{"id": id})
}
