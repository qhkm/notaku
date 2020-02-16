package routes

import (
	"github.com/gin-gonic/gin"
	"notaku/controller"
)

// Routes this is routes
func Routes(route *gin.Engine) {
	v1 := route.Group("api/v1")
	{
		// posts
		v1.GET("/post/:id", controller.GetPost)
		v1.GET("/posts", controller.GetAllPosts)
		v1.POST("/posts/add", controller.AddPost)
		v1.PUT("/post/:id", controller.UpdatePost)
		v1.DELETE("/post/:id", controller.DeletePost)

		// users
		// PRIVATE
		v1.GET("/user/:id", controller.GetUser)
		v1.GET("/users", controller.GetAllUsers)
		v1.POST("/users/add", controller.AddUser)
		v1.PUT("/user/:id", controller.UpdateUser)
		v1.DELETE("/user/:id", controller.DeleteUser)

		// auth
		v1.POST("/login", controller.Login)
		v1.GET("/signup", controller.Signup)
		v1.GET("/logout", controller.Logout)
	}
}
