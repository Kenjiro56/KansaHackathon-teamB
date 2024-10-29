package controller

import (
	"KansaiHack-Friday/db"
	"KansaiHack-Friday/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateObj(c *gin.Context) { //objの作成
	// 追加する際にuserが存在するかどうかのチェックを入れたい！！
	var obj models.Obj
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&obj).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// default
func GetAll(c *gin.Context) {
	var objs []models.Obj
	if err := db.DB.Find(&objs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
		return
	}
	c.JSON(http.StatusOK, objs)
}

// preload
// func GetSingleObj(c *gin.Context) {

// }
