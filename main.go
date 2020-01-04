package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"notaku/controller"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api := r.Group("api")
	{
		api.GET("/post/id", controller.GetPost)
		api.GET("/posts", controller.GetAllPosts)
		api.GET("/posts/add", controller.AddPost)
		api.GET("/posts/delete", controller.DeletePost)
	}
	r.Run()
}

//check error
func checkError(err error) {
	if err != nil {
		panic(err)
	}

}

// not found page
func NotFound(route *gin.Engine) {
	route.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(404, "Not Found")
	})
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
