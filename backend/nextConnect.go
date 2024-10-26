package nextConnect

import (
	"github.com/gin-gonic/gin"
)

// cors設定
router.Use(func(cors *gin.Context) {
		cors.Header("Access-Control-Allow-Origin", "http://localhost:8080") // 一旦ローカルホストに設定(本番は環境変数で設定するのが良さそう)
		cors.Writer("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS") // 一旦GET, POST, DELETEのみ許可
		cors.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		cors.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if cors.Request.Method == "OPTIONS" {
			cors.AbortWithStatus(200)
			return
		}
		cors.Next()
})
