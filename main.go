package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

type Posts struct {
	Items []Post
}
type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api := r.Group("api")
	{
		api.GET("/", getPong)
		api.GET("/test", getTest)
		api.GET("/ayam", getPong)
		api.GET("/itik", getPong)
		api.GET("/udang", getPong)
	}
	r.Run()
}

func getPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
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

// write code to open file
