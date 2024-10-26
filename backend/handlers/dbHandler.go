package handlers

import (
	"net/http"

	"KansaiHack-Friday/db"

	"github.com/gin-gonic/gin"
)

func CheckDBConnection(c *gin.Context) {
	err := db.DB.Exec("SELECT 1").Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Database connection failed", "error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Database connection is healthy"})
	}
}
