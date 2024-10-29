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
		auth.GET("/allusers", controller.GetUsers)
		auth.GET("/user/:id", controller.GetUser)
		auth.DELETE("/deleteuser/:id", controller.DeleteUser)
		auth.GET("/select", nextConnect.TestNext)
		auth.POST("/insert", nextConnect.TestNext2)
	}
}