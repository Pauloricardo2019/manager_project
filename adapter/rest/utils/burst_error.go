package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BurstError(c *gin.Context, err error, statusCode int) {
	switch {
	case statusCode == 400:
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
	case statusCode == 404:
		c.JSON(http.StatusNotFound, gin.H{"error ": err.Error()})
	case statusCode == 401:
		c.JSON(http.StatusUnauthorized, gin.H{"error ": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error ": err.Error()})
	}
	return
}
