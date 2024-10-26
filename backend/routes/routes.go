package routes

import (
	"KansaiHack-Friday/controller"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(r gin.IRouter) {

	auth := r.Group("/")
	{
		auth.POST("/signup", controller.CreateUser)
		auth.GET("/allusers", controller.GetUsers)
		auth.GET("/user/:id", controller.GetUser)
	}
}
