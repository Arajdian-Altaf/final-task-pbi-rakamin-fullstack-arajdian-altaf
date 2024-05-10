package router

import (
	"github.com/Arajdian-Altaf/final-task-pbi/controllers"
	"github.com/Arajdian-Altaf/final-task-pbi/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
	user := route.Group("/users") 
	{
		user.POST("/register", controllers.UserCreate)
		user.GET("/login", controllers.UserLogin)
		user.PUT("/:userId", middlewares.JWTMiddleware(), controllers.UserUpdate)
		user.DELETE("/:userId", middlewares.JWTMiddleware(), controllers.UserDelete)
	}
}