package routes

import (
	"KansaiHack-Friday/db"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(r gin.IRouter) {

	database := r.Group("/db")
	{
		database.GET("/", db.Connect)
	}
}
