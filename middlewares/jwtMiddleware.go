package middlewares

import (
	"net/http"

	"github.com/Arajdian-Altaf/final-task-pbi/helpers"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {		
		tokenString := strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", "")

		userClaims, err := helpers.ParseToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Set("userClaims", userClaims)
		c.Next()
	}
}	
