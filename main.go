package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Post struct {
	id int
	Title string
	Body string
}

type Posts struct {
	Items []Post
}

// create new article
func New() *Post {
	post1 := Post{
		id: 1,
		Title: "test",
		Body: "body",
	}

	return &post1
}

func (p *Post) getAllPosts() {
	return p.id
}

func PostsGet(post *Post) gin.HandlerFunc {
	return func(c *gin.Context){
		result := post.getAllPosts()
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
