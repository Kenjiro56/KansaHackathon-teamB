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

		// ユーザーからの入力を取得
		userInput := c.Query("input") // クエリパラメータから取得

		// 入力がなかった場合のエラーハンドリング
		if userInput == "" {
			c.JSON(400, gin.H{"error": "入力が必要です。"})
			return
		}

		// メッセージの例
		// messages := []ChatMessage{
		// 	{Role: "system", Content: "ユーザーが「〇〇がしたい」または「〇〇が達成したい」などの入力に従い、あなたはその目標を達成するための具体的な道筋を示してください。目標に応じて、必要な定量的なToDoを生成し、リスト形式で提示します。リストには各ToDoに対する簡単な説明を含め、ユーザーが次に何をすべきかを明確に理解できるようにしてください。"}, // システムメッセージで言語指定
		// 	{Role: "user", Content: "英会話が上達したい"},
		// }

		// 入力メッセージの準備
		messages := []ChatMessage{
			{Role: "system", Content: "あなたは日本語を流暢に話すAIです。"}, // システムメッセージで言語指定
			{Role: "user", Content: userInput},
		}

		// OpenAI APIへのリクエスト
		response, err := OpenAIAPI("gpt-3.5-turbo", messages) // 適切なモデルを指定

		// エラーハンドリング
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		// 正常時
		c.JSON(200, gin.H{"response": response})
	})

	// サーバー起動
	router.Run(":8080")
}
