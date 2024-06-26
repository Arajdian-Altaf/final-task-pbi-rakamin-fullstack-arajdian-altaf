package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/Arajdian-Altaf/final-task-pbi/helpers"
	"github.com/Arajdian-Altaf/final-task-pbi/models"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func UserCreate(c *gin.Context) {
	DB := c.MustGet("db").(*gorm.DB)

	var userBody struct {
		Username string `valid:"required,ascii"`
		Email    string `valid:"required,email"`
		Password string `valid:"required,minstringlength(6)"`
	}

	err := c.BindJSON(&userBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty body",
		})
		return
	}

	if _, err := valid.ValidateStruct(userBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, err := helpers.HashPassword(userBody.Password)

	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{Username: userBody.Username, Email: userBody.Email, Password: hashedPassword}

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

func UserLogin(c *gin.Context) {
	DB := c.MustGet("db").(*gorm.DB)

	var loginRequest struct {
		Email    string `valid:"email"`
		Password string `valid:"minstringlength(6)"`
	}

	err := c.BindJSON(&loginRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty body",
		})
		return
	}

	if _, err := valid.ValidateStruct(loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User

	result := DB.First(&user, "email = ?", loginRequest.Email)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "User not found",
		})
		return
	}

	if !helpers.CheckPassword(loginRequest.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Wrong password",
		})
		return
	}

	userClaim := helpers.UserClaims{
		ID:             user.ID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	accessToken, err := helpers.GenerateToken(userClaim)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": accessToken,
	})
}

func UserUpdate(c *gin.Context) {
	DB := c.MustGet("db").(*gorm.DB)
	userClaims := c.MustGet("userClaims").(*helpers.UserClaims)

	var userBody struct {
		Username string
		Email    string `valid:"email"`
		Password string `valid:"minstringlength(6)"`
	}

	err := c.BindJSON(&userBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty body",
		})
		return
	}
	
	if _, err := valid.ValidateStruct(userBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Get id from URL
	id := c.Param("userId")

	var user models.User
	result := DB.First(&user, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	if user.ID != userClaims.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Cannot update other user",
		})
		return
	}

	hashedPassword, err := helpers.HashPassword(userBody.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	DB.Model(&user).Updates(models.User{
		Username: userBody.Username,
		Email:    userBody.Email,
		Password: hashedPassword,
	})

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UserDelete(c *gin.Context) {
	userClaims := c.MustGet("userClaims").(*helpers.UserClaims)
	DB := c.MustGet("db").(*gorm.DB)

	// Get id from URL
	id := c.Param("userId")

	var user models.User
	result := DB.First(&user, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	if user.ID != userClaims.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Cannot delete other user",
		})
		return
	}

	result = DB.Select("Photos").Delete(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}