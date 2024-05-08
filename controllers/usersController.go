package controllers

import (
	"net/http"
	"github.com/Arajdian-Altaf/final-task-pbi/models"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func UserCreate(c *gin.Context) {
	DB := c.MustGet("db").(*gorm.DB)

	var userBody struct {
		Username string
		Email    string `valid:"email"`
		Password string `valid:"minstringlength(6)"`
	}

	c.Bind(&userBody)

	if _, err := valid.ValidateStruct(userBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{Username: userBody.Username, Email: userBody.Email, Password: userBody.Password}

	result := DB.Create(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}
