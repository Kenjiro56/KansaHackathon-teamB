package controller

import (
	"KansaiHack-Friday/db"
	"KansaiHack-Friday/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateTodo - 新しいTodoを作成する
func CreateTodo(c *gin.Context) {
	var todo models.Todo

	// リクエストボディからデータをバインド
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	// データベースに新しいTodoを作成
	if err := db.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// GetTodo - 指定されたIDのTodoを取得する(Preload)
func GetTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var todo models.Todo
	// IDに基づいてTodoを取得
	if err := db.DB.Preload("User").Preload("Obj").First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// GetTodos - すべてのTodoを取得する(Preload)
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	// すべてのTodoを取得
	if err := db.DB.Preload("User").Preload("Obj").Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// UpdateTodo - 指定されたIDのTodoを更新する
func UpdateTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var todo models.Todo
	// Todoが存在するか確認
	if err := db.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// リクエストボディから更新内容を取得
	var todoUpdate models.Todo
	if err := c.ShouldBindJSON(&todoUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// フィールドの更新
	todo.Status = todoUpdate.Status
	todo.Progress = todoUpdate.Progress
	todo.MaxProgress = todoUpdate.MaxProgress
	todo.DeleteFlag = todoUpdate.DeleteFlag
	todo.UpdatedAt = time.Now()

	// データベースに保存
	if err := db.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DeleteTodo - 指定されたIDのTodoを削除する
func DeleteTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var todo models.Todo
	// Todoが存在するか確認
	if err := db.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// データベースから削除
	if err := db.DB.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
