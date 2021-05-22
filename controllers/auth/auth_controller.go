package auth

import (
	"go-auth/controllers/services"
	dto "go-auth/dto/login"
	"go-auth/repositories/user"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req dto.LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request body",
		})
		return
	}

	repo := user.UserRepository{}
	service := services.NewLoginService(repo)

	token, err := service.GetToken(req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func Logout(c *gin.Context) {
	c.Status(204)
}
