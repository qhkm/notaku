package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"notaku/model"
	"notaku/util"
	//comment
	_ "github.com/mattn/go-sqlite3"
)

// GetPost test
func GetPost(c *gin.Context) {
	db, err := model.NewDB()
	ut := util.Util{}
	ut.CheckError(err)
	postService := model.Env{DB: db.DB}
	PostList := postService.GetPostList()
	// post := model.Post{}
	// err := c.Bind(&post)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, "Bad request")
	// 	return
	// }
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    PostList,
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
