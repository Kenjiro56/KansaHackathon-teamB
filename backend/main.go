package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//Ginフレームワークのデフォルトの設定を使用してルータを作成
	router := gin.Default()

	// ルートハンドラの定義
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!!!waiwai",
		})
	})

	router.GET("/openai", func(c *gin.Context) {
		messages := []ChatMessage{
			{Role: "user", Content: "Tell me a joke."},
		}
		response, err := OpenAIAPI("gpt-3.5-turbo", messages) // 適切なモデルを指定
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"response": response})
	})

	// サーバー起動
	router.Run(":8080")
}
