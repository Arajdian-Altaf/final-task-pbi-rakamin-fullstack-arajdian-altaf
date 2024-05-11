package router

import (
	"github.com/Arajdian-Altaf/final-task-pbi/controllers"
	"github.com/Arajdian-Altaf/final-task-pbi/middlewares"
	"github.com/gin-gonic/gin"
)

func PhotoRoutes(route *gin.Engine) {
	photo := route.Group("/photos")
	{
		photo.GET("/", middlewares.JWTMiddleware(), controllers.PhotoGet)
		photo.POST("/", middlewares.JWTMiddleware(), controllers.PhotoCreate)
		photo.PUT("/:photoId", middlewares.JWTMiddleware(), controllers.PhotoUpdate)
		photo.DELETE("/:photoId", middlewares.JWTMiddleware(), controllers.PhotoDelete)
	}
}
