package controller

import (
	"KansaiHack-Friday/db"
	"KansaiHack-Friday/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateObjInput struct {
	ObjTitle   *string `json:"obj_title"`
	Progress   *int    `json:"progress"`
	DeleteFlag *bool   `json:"delete_flag"`
}

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
	c.JSON(http.StatusOK, gin.H{
		"message": "Object created successfully",
		"data":    obj, // 作成されたオブジェクトをレスポンスとして含める
	})
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
	id := c.Param("id")

	// URLパラメータで指定されたIDのObjを取得
	var obj models.Obj
	if err := db.DB.First(&obj, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Obj not found"})
		return
	}

	// リクエストボディから更新データを取得
	var input UpdateObjInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// フィールドごとに存在するかどうかを確認し、存在する場合のみ更新
	if input.ObjTitle != nil {
		obj.ObjTitle = *input.ObjTitle
	}
	if input.Progress != nil {
		obj.Progress = *input.Progress
	}
	if input.DeleteFlag != nil {
		obj.DeleteFlag = *input.DeleteFlag
	}

	// UpdatedAt フィールドを現在の時間に更新
	obj.UpdatedAt = time.Now()

	// データベースに保存
	if err := db.DB.Save(&obj).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update obj"})
		return
	}

	c.JSON(http.StatusOK, obj)
}
