package auth

import (
	"notaku/controller"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	v1 := route.Group("api/v1")
	{
		v1.GET("/post/:id", controller.GetPost)
		v1.GET("/posts", controller.GetAllPosts)
		v1.POST("/posts/add", controller.AddPost)
		v1.PUT("/post/:id", controller.UpdatePost)
		v1.DELETE("/post/:id", controller.DeletePost)
		v1.GET("/login", controller.Login)
		v1.GET("/signup", controller.Signup)
		v1.GET("/logout", controller.Logout)
	}
}