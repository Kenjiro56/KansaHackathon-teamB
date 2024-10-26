package routes

import (
	"KansaiHack-Friday/handlers"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(r gin.IRouter) {

	database := r.Group("/db")
	{
		database.GET("/", handlers.CheckDBConnection)
	}
}
