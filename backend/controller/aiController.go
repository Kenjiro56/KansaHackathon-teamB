package controller

import (
	"KansaiHack-Friday/APIConnect"
	"net/http"
	"strings"

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
		{Role: "system",
			Content: "あなたはユーザーの目標に対するTodoタスクを提案するアシスタントです。各タスクは簡潔に記述してください。例:「毎日30分運動する」のように。"},
		{Role: "user", Content: goal + "を達成したいです。Todoタスクを番号なしで羅列して提案してください。"},
	}

	response, err := APIConnect.OpenAIAPI("gpt-3.5-turbo", messages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// AIの応答を分割してリスト形式に整形
	tasks := parseResponseToTasks(response)

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
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
		{Role: "user", Content: consultation + "という相談についてToDoタスクを作成し直してください。"},
	}

	response, err := APIConnect.OpenAIAPI("gpt-3.5-turbo", messages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// AIの応答を分割してリスト形式に整形
	tasks := parseResponseToTasks(response)

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// parseResponseToTasks はAIの応答をリスト形式に変換する
func parseResponseToTasks(response string) []string {
	// 応答を改行で分割してリストを作成
	tasks := strings.Split(response, "\n")
	// 空のタスクをフィルタリング
	var filteredTasks []string
	for _, task := range tasks {
		task = strings.TrimSpace(task) // 前後の空白を削除
		if task != "" {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}
