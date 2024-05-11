package controllers

import (
	"net/http"

	"github.com/Arajdian-Altaf/final-task-pbi/helpers"
	"github.com/Arajdian-Altaf/final-task-pbi/models"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PhotoCreate(c *gin.Context) {
	DB := c.MustGet("db").(*gorm.DB)
	userClaims := c.MustGet("userClaims").(*helpers.UserClaims)

	var photoBody struct {
		Title    string `valid:"required,ascii"`
		Caption  string `valid:"required,ascii"`
		PhotoURL string `json:"photo_url" valid:"required,url"`
	}

	err := c.BindJSON(&photoBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty body",
		})
		return
	}

	if _, err := valid.ValidateStruct(photoBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	result := DB.First(&user, userClaims.ID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	photo := models.Photo{Title: photoBody.Title, Caption: photoBody.Caption, PhotoURL: photoBody.PhotoURL, UserID: user.ID}

	result = DB.Create(&photo)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
	}

	c.Status(http.StatusCreated)
}
