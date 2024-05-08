package router

import "github.com/gin-gonic/gin"

func PhotoRoutes(route *gin.Engine) {
	photo := route.Group("/photos")
	{
		photo.POST("/", func(c *gin.Context) {})
		photo.GET("/", func(c *gin.Context) {})
		photo.PUT("/:photoId", func(c *gin.Context) {})
		photo.DELETE("/:photoId", func(c *gin.Context) {})
	}
}
