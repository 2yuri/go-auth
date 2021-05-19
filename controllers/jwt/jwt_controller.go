package jwt

import "github.com/gin-gonic/gin"

func ValidToken(c *gin.Context) {
	c.Status(204)
}
