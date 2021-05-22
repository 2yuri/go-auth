package jwt

import (
	"go-auth/services"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ValidToken(c *gin.Context) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	tokenSrting := authHeader[len(BEARER_SCHEMA):]

	log.Print(tokenSrting)

	token, err := services.NewJWTService().ValidateToken(tokenSrting)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	if token.Valid {
		c.JSON(200, gin.H{
			"token": token.Claims.(jwt.MapClaims),
		})
	}

}
