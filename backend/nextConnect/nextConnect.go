package nextConnect

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func NextConnect() *gin.Engine {
	connector := gin.Default()

	// cors設定
	connector.Use(func(cors *gin.Context) {
		cors.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")       // 一旦ローカルホストに設定(本番は環境変数で設定するのが良さそう)
		cors.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS") // 一旦GET, POST, DELETEのみ許可
		cors.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		cors.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if cors.Request.Method == "OPTIONS" {
			cors.AbortWithStatus(200)
			return
		}
		cors.Next()
	})

	// 上記は一旦放置（nextからapiを叩けるようにする）

	// test用のエンドポイント
	// jsonの参照
	connector.GET("/select", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!!!waiwai2",
		})
		fmt.Println("Hello, World!!!waiwai2")
	})

	// 追加用のエンドポイント
	connector.POST("/insert", func(ctx *gin.Context) {
		type Message struct {
			Message string `json:"message"`
		}

		var req Message
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{
				"message": "Bad Request",
			})
			return
		}

		// ここにデータベース追加系書くとグッド
		fmt.Println(req.Message)

		ctx.JSON(201, "Success")
	})

	return connector
}
