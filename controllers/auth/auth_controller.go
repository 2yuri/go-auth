package auth

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	token := "a2hsash1h23asjd"
	c.JSON(200, gin.H{
		"token": token,
	})
}

func Logout(c *gin.Context) {
	c.Status(204)
}
