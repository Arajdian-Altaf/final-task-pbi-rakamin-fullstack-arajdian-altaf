package router

import "github.com/gin-gonic/gin"

func UserRoutes(route *gin.Engine) {
	user := route.Group("/users") 
	{
		user.POST("/register", func(c *gin.Context) {})
		user.GET("/login", func(c *gin.Context) {})
		user.PUT("/:userId", func(c *gin.Context) {})
		user.DELETE("/:userId", func(c *gin.Context) {})
	}
}