package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"notaku/config"
	"notaku/model"
	"notaku/util"
	//comment
	_ "github.com/mattn/go-sqlite3"
)

var db *config.DB

func init() {
	var err error
	db, err = config.NewDB()
	ut := util.Util{}
	ut.CheckError(err)
}

// GetPost test
func GetPost(c *gin.Context) {
	pathParam := c.Param("id")
	id, _ := strconv.Atoi(pathParam)

	postService := model.Env{DB: db.DB}
	PostList := postService.SinglePost(id)
	c.JSONP(http.StatusOK, gin.H{
		"message": "get all post",
		"data":    PostList,
		"param":   pathParam,
	})
}

// GetAllPosts tets
func GetAllPosts(c *gin.Context) {
	postService := model.Env{DB: db.DB}
	PostList := postService.AllPosts()
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

// UpdatePost test
func UpdatePost(c *gin.Context) {
	pathParam := c.Param("id")
	id, _ := strconv.Atoi(pathParam)
	objA := model.Post{}
	if errA := c.ShouldBind(&objA); errA != nil {
		panic(errA.Error())
	}
	fmt.Println(objA)
	objA.ID = id
	fmt.Println(objA)
	postService := model.Env{DB: db.DB}
	ok, err := postService.UpdatePost(objA)

	if err != nil {
		c.JSON(400, gin.H{
			"error":  err.Error(),
			"status": ok,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success update post",
	})
}

// AddPost test
func AddPost(c *gin.Context) {
	objA := model.Post{}
	if errA := c.ShouldBind(&objA); errA != nil {
		panic(errA.Error())
	}
	postService := model.Env{DB: db.DB}

	ok, err := postService.AddPost(objA)
	if err != nil {
		c.JSON(400, gin.H{
			"error":  err.Error(),
			"status": ok,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success add post",
	})
}

// DeletePost test
func DeletePost(c *gin.Context) {
	pathParam := c.Param("id")
	id, _ := strconv.Atoi(pathParam)
	postService := model.Env{DB: db.DB}
	// get id from param
	// id := c.Params.ByName("id")
	rowsAffected, err2 := postService.DeletePost(id)
	if err2 != nil {
		panic(err2.Error())
	}
	fmt.Println(rowsAffected)
	c.JSON(http.StatusOK, gin.H{
		"message": "delete OK",
	})
}
