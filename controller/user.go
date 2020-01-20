package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"notaku/model"
	//comment
	_ "github.com/mattn/go-sqlite3"
)

// GetAllPosts tets
func GetAllUsers(c *gin.Context) {
	userService := model.Env{DB: db.DB}
	UserList := userService.AllUsers()
	// post := model.Post{}
	// err := c.Bind(&post)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, "Bad request")
	// 	return
	// }
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    UserList,
	})

}