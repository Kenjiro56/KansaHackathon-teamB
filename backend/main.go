package main

import (
	"KansaiHack-Friday/nextConnect"
)

func main() {
	router := nextConnect.NextConnect()
	//Ginフレームワークのデフォルトの設定を使用してルータを作成
	// router := gin.Default()

	// // ルートハンドラの定義
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello, World!!!waiwai2",
	// 	})
	// })

	// サーバー起動
	router.Run(":8080")
}
