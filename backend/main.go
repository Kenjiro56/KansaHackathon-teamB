package main

import (
	"KansaiHack-Friday/db"
	"KansaiHack-Friday/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	//DBに接続
	if err := db.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Ginフレームワークのデフォルトの設定を使用してルータを作成
	router := gin.Default()
	routes.DefineRoutes(router)

	// routingのログを出す
	for _, route := range router.Routes() {
		fmt.Printf("Method: %s, Path: %s\n", route.Method, route.Path)
	}

	// サーバー起動
	router.Run(":8080")
}
