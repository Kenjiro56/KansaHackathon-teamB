package routes

import (
	"KansaiHack-Friday/controller"
	"KansaiHack-Friday/nextConnect"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(r gin.IRouter) {

	auth := r.Group("/")
	{
		auth.POST("/signup", controller.CreateUser)
		auth.POST("/login", controller.Login)
		auth.GET("/allusers", controller.GetUsers)
		auth.GET("/user/:id", controller.GetUser)
		auth.DELETE("/deleteuser/:id", controller.DeleteUser)
		auth.GET("/select", nextConnect.TestNext)
		auth.POST("/insert", nextConnect.TestNext2)
	}

	obj := r.Group("/obj")
	{
		obj.POST("/add", controller.CreateObj)
		obj.GET("/getAll", controller.GetAll)
		obj.GET(":id", controller.GetSingleObj)
		obj.DELETE("/deleteObj/:id", controller.DeleteObj)
		obj.PUT("/update/:id", controller.UpdateObj)
	}

	// AI関連のエンドポイント
	ai := r.Group("/ai")
	{
		// 目標に対するフィードバック
		ai.GET("/goal", controller.GenerateGoalTasks)
		// 相談内容に対するフィードバック
		ai.GET("/consultation", controller.ConsultationFeedback)
	}

	todo := r.Group("/todo")
	{
		todo.POST("/", controller.CreateTodo)
		todo.GET("/getAll", controller.GetTodos)
		todo.GET("/:id", controller.GetTodo)
		todo.PUT("/update/:id", controller.UpdateTodo)
		todo.DELETE("/delete/:id", controller.DeleteTodo)
	}
}
