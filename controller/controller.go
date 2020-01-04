package controller

import (
	"net/http"

	"notaku/model"

	"github.com/gin-gonic/gin"
	//comment
	_ "github.com/mattn/go-sqlite3"
)

// GetPost test
func GetPost(c *gin.Context) {
	post := model.Post{}
	err := c.Bind(&post)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "get a post",
	})
}

// GetAllPosts tets
func GetAllPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get all post",
	})
}

// AddPost test
func AddPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "add a post",
	})
}

// DeletePost test
func DeletePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "delete post",
	})
}
