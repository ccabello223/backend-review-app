package visited

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetVisitedCountController(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	count, err := GetVisitedCountService(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"visited_count": count})
}

func GetVisitorsCount(c *gin.Context) {
	placeIDStr := c.Param("place_id")
	placeID, err := strconv.Atoi(placeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid place ID"})
		return
	}

	count, err := GetVisitorsCountService(uint(placeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get visitor count"}) // Don't expose internal errors directly in production.
		return
	}

	c.JSON(http.StatusOK, gin.H{"visitors_count": count})
}

func CreateVisitedPlace(c *gin.Context) {
	var input VisitedPlaceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID"})
		return
	}

	if err := CreateVisitedPlaceService(userId.(uint), input.PlaceId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Visited place created successfully"})
}

func DeleteVisitedPlace(c *gin.Context) {
	var input VisitedPlaceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID"})
		return
	}

	if err := DeleteVisitedPlaceService(userId.(uint), input.PlaceId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Visited place deleted successfully"})

}
