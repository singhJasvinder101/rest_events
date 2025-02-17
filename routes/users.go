package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/singhJasvinder101/models"
	"github.com/singhJasvinder101/utils"
)

func signup(c *gin.Context){
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Please enter the required fields correctly",
		})
		return
	}

	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	user.Password = hashedPassword
	err = user.Save()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(201, gin.H{
		"message": "User created successfully",
		"user": user,
	})
}

func login(c *gin.Context){
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Please enter the required fields correctly",
		})
		return
	}

	err = user.Authenticate()

	if err != nil {
		c.JSON(401, gin.H{
			"message": "could not authenticate user",
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User logged in successfully",
		"user": token,
	})
}

