package main

import (
	"KansaiHack-Friday/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//Ginフレームワークのデフォルトの設定を使用してルータを作成
	router := gin.Default()

	// ルート定義を呼び出す
	routes.DefineRoutes(router)

	// ルートハンドラの定義
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!!!waiwai",
		})
	})

	// サーバー起動
	router.Run(":8080")
}
