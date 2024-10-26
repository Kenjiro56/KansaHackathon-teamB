package main

import (
	"KansaiHack-Friday/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Ginフレームワークのデフォルトの設定を使用してルータを作成
	router := gin.Default()

	// ルートハンドラの定義
	router.GET("/", func(c *gin.Context) {
		fmt.Printf("-----------------------------------")
		c.JSON(200, gin.H{
			"message": "Hello, World!!!waiwai22",
		})
	})

	router.GET("/dbConnect", func(c *gin.Context) {
		db.Connect(c)
	})

	for _, route := range router.Routes() {
		fmt.Printf("-----------------------------------123")
		fmt.Printf("Method: %s, Path: %s\n", route.Method, route.Path)
	}

	// サーバー起動
	router.Run(":8080")
}
