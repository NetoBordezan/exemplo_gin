package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy", "timestamp": time.Now(), "version": "1.0.0"})
}
