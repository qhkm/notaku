package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"notaku/controller"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3500"}
	r.Use(cors.New(config))
	r.NoRoute(notFound)
	api := r.Group("api")
	{
		v1 := api.Group("v1")
		v1.GET("/post/:id", controller.GetPost)
		v1.GET("/posts", controller.GetAllPosts)
		v1.POST("/posts/add", controller.AddPost)
		v1.PUT("/post/:id", controller.UpdatePost)
		v1.DELETE("/post/:id", controller.DeletePost)
		v1.GET("/login", controller.Login)
		v1.GET("/signup", controller.Signup)
		v1.GET("/logout", controller.Logout)
	}
	r.Run()
}

// not found page
func notFound(c *gin.Context) {
	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}

// nomethods commentj
func NoMethods(route *gin.Engine) {
	route.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(405, "not allowed")
	})
}

type QueryError struct {
	Query string
	Err   error
}
