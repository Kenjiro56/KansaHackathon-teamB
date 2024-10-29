package main

import (
	"KansaiHack-Friday/db"
	"KansaiHack-Friday/routes"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//DBに接続
	if err := db.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// マイグレーションを実行
	if err := db.Migrate(); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Ginフレームワークのデフォルトの設定を使用してルータを作成
	router := gin.Default()

	// CORSの設定
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000/"}, // 許可するオリジン
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routes.DefineRoutes(router)

	// routingのログを出す
	for _, route := range router.Routes() {
		fmt.Printf("Method: %s, Path: %s\n", route.Method, route.Path)
	}

	// サーバー起動
	router.Run(":8080")
}
