package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type Blog struct {
	id int
	title string
	body string
}

var blogList = []Blog{
	Blog{
		id: 1, title: "test1", body: "testbody1",
	},
	Blog{
		id: 2, title: "test2", body: "testbody2",
	},
}

// return array of blog
func getBlog() Blog {
	firstBlog := Blog{
		id: 1,
		title: "testing title",
		body: "testing body",
	}
	return firstBlog
}

func getBlogList() []Blog {
	return blogList
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	blog := getBlog()
	blogList := getBlogList()
	fmt.Println(blogList)
	r.GET("/ping", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": blog,
		})
	})
	r.Run()
}
