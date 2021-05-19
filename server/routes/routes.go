package routes

import (
	"go-auth/controllers/auth"
	"go-auth/controllers/jwt"
	"go-auth/controllers/user"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(engine *gin.Engine) *gin.Engine {
	main := engine.Group("api/v1")
	{
		user_g := main.Group("user")
		{
			user_g.POST("/", user.CreateUser)
			user_g.PUT("/", user.EditUser)
			user_g.DELETE("/:id", user.DeleteUser)
		}

		jwt_g := main.Group("jwt")
		{
			jwt_g.GET("valid", jwt.ValidToken)
		}

		main.POST("/login", auth.Login)
		main.GET("/logout", auth.Logout)
	}

	return engine
}
