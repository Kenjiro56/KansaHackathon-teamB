package main

import (
	"github.com/gin-gonic/gin"
)
func NextConnect(c *gin.Context) {
    // 関数の実装
    c.JSON(200, gin.H{
        "message": "Next Connect Successful",
    })
}
