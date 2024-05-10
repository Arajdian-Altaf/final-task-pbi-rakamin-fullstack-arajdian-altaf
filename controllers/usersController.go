package controllers

import (
	"log"
	"net/http"

	"github.com/Arajdian-Altaf/final-task-pbi/helpers"
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

	loginRequest.Email = c.Query("email")
	loginRequest.Password = c.Query("password")

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

	c.Bind(&userBody)

	if _, err := valid.ValidateStruct(userBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	DB.Model(&user).Updates(models.User{
		Username: userBody.Username,
		Email: userBody.Email,
		Password: userBody.Password,
	})

	c.JSON(200, gin.H{
		"user": user,
	})
}
