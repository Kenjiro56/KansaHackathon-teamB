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
		auth.DELETE("/deleteuser/:id", controller.DeleteUser)
	}

	obj := r.Group("/obj")
	{
		obj.POST("/add", controller.CreateObj)
		obj.GET("/getAll", controller.GetAll)
	}
}
