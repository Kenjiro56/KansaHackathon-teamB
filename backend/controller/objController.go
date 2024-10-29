package controller

import (
	"KansaiHack-Friday/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateObj(c *gin.Context) { //objの作成
	var obj models.Obj
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
