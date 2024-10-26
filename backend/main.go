package main

import (
	"net/http"

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

	// ルートハンドラの定義
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil) // index.htmlというテンプレートをレンダリング
	})
	// テンプレートの設定
	router.LoadHTMLGlob("templates/*") // templatesディレクトリ内のHTMLファイルを読み込む

	router.GET("/openai", func(c *gin.Context) {

		// ユーザーからの入力を取得
		userInput := c.Query("input") // クエリパラメータから取得

		// 入力がなかった場合のエラーハンドリング
		if userInput == "" {
			c.JSON(400, gin.H{"error": "入力が必要です。"})
			return
		}

		// メッセージの準備
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
