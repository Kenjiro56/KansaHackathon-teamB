package routes

import (
	"KansaiHack-Friday/controller"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(r gin.IRouter) {

	// AI関連のエンドポイント
	ai := r.Group("/ai")
	{
		// 目標に対するフィードバック
		ai.GET("/goal", controller.GenerateGoalTasks)
		// 相談内容に対するフィードバック
		ai.GET("/consultation", controller.ConsultationFeedback)
	}
}
