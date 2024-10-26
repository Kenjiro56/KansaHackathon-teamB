package routes

import (
	"KansaiHack-Friday/controller"
	"KansaiHack-Friday/handlers"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(r gin.IRouter) {

	database := r.Group("/db")
	{
		database.GET("/", handlers.CheckDBConnection)
	}
	auth := r.Group("/")
	{
		auth.POST("/signup", controller.CreateUser)
	}
}
