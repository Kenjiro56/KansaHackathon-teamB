package controller

import (
	"KansaiHack-Friday/db"
	"KansaiHack-Friday/models"
	"net/http"
	"strconv"

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
func GetSingleObj(c *gin.Context) {
	id := c.Param("id")
	var obj models.Obj
	if err := db.DB.First(&obj, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Obj not found"})
		return
	}
	c.JSON(http.StatusOK, obj)
}

func DeleteObj(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Obj{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// UpdateObj は Obj データを更新するコントローラーです
func UpdateObj(c *gin.Context) {
	// URLからIDを取得して、uintに変換
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// 更新内容をリクエストから取得
	var objUpdate models.Obj
	if err := c.ShouldBindJSON(&objUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// DBから対象のObjを検索
	var obj models.Obj
	if err := db.DB.First(&obj, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Obj not found"})
		return
	}

	// 更新内容を設定
	obj.ObjTitle = objUpdate.ObjTitle
	obj.Progress = objUpdate.Progress
	obj.DeleteFlag = objUpdate.DeleteFlag

	// データベースで更新
	if err := db.DB.Save(&obj).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Obj"})
		return
	}

	c.JSON(http.StatusOK, obj)
}
