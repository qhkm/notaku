package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

type Posts struct {
	ID    int
	Items []Post
}

// create new article
func New() []Post {
	lol := []Post{
			{Id: 1, Title: "test1", Body: "test1"},
			{Id: 2, Title: "test2", Body: "test2"},
		}
	return lol
}

func (p *[]Post) Add(post []Post) Posts {
	newPost := 
		Post{
			Id:    1,
			Title: "Add more",
			Body:  "More body",
		},

	p = append(p, newPost)
	return p

	newAddPost := append(newPost, basicPost)
	return newAddPost
}

// get all posts
func (p Posts) getAllPosts() Posts {
	return p
}

// get all post
func PostsGet(posts Posts) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := posts.getAllPosts()
		c.JSON(http.StatusOK, result)
	}
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	posts := New()
	r.GET("/ping", PostsGet(posts))
	r.Run()
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
