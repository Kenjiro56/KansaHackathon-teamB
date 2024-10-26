package db

import "github.com/gin-gonic/gin"

func Connect(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DBConnect Successful",
	})
}
