package controller

import (
	"KansaiHack-Friday/APIConnect"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GoalFeedback は目標に対するフィードバックを返すエンドポイント
func GenerateGoalTasks(c *gin.Context) {
	goal := c.Query("goal")
	if goal == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "目標が入力されていません。"})
		return
	}

	// AIからの応答を生成するためのロジック
	messages := []APIConnect.ChatMessage{
		{Role: "system", Content: "あなたはユーザーの目標に対するフィードバックとTodoタスクを提案するアシスタントです。"},
		{Role: "user", Content: goal + "を達成したいです。アドバイスとTodoタスクをください。"},
	}

	response, err := APIConnect.OpenAIAPI("gpt-3.5-turbo", messages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": response})
}

// ConsultationFeedback は相談内容に対するフィードバックを返すエンドポイント
func ConsultationFeedback(c *gin.Context) {
	consultation := c.Query("consultation")
	if consultation == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "相談内容が入力されていません。"})
		return
	}

	// AIからの応答を生成するためのロジック
	messages := []APIConnect.ChatMessage{
		{Role: "system", Content: "あなたはユーザーのTodo進捗における相談に答えるアシスタントです。"},
		{Role: "user", Content: consultation + "という相談について助言をください。"},
	}

	response, err := APIConnect.OpenAIAPI("gpt-3.5-turbo", messages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": response})
}
