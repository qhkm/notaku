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
	ID int
	Items []Post
}

// create new article
func New() Posts {
	p := Posts {
		ID: 1,
		Items: []Post {
			{id: 1, Title: "test1", Body: "test1"},
			{id: 2, Title: "test2", Body: "test2"},
		},
	}
	return p
}
// get all posts
func (p Posts) getAllPosts() Posts {
	return p
}

func PostsGet(posts Posts) gin.HandlerFunc {
	return func(c *gin.Context){
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
