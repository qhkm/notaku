package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"notaku/model"
	//comment
	_ "github.com/mattn/go-sqlite3"
)


// GetUser test
func GetUser(c *gin.Context) {
	pathParam := c.Param("id")
	id, _ := strconv.Atoi(pathParam)

	userService := model.Env{DB: db.DB}
	UserList := userService.SingleUser(id)
	c.JSONP(http.StatusOK, gin.H{
		"message": "get one user",
		"data":    UserList,
		"param":   pathParam,
	})
}

// GetAllUsers tets
func GetAllUsers(c *gin.Context) {
	userService := model.Env{DB: db.DB}
	UserList := userService.AllUsers()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    UserList,
	})

}


// UpdateUser test
func UpdateUser(c *gin.Context) {
	pathParam := c.Param("id")
	id, _ := strconv.Atoi(pathParam)
	objA := model.User{}
	if errA := c.ShouldBind(&objA); errA != nil {
		panic(errA.Error())
	}
	fmt.Println(objA)
	objA.ID = id
	fmt.Println(objA)
	userService := model.Env{DB: db.DB}
	ok, err := userService.UpdateUser(objA)

	if err != nil {
		c.JSON(400, gin.H{
			"error":  err.Error(),
			"status": ok,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success update user",
	})
}

// AddUser test
func AddUser(c *gin.Context) {
	objA := model.User{}
	if errA := c.ShouldBind(&objA); errA != nil {
		panic(errA.Error())
	}
	userService := model.Env{DB: db.DB}

	ok, err := userService.AddUser(objA)
	if err != nil {
		c.JSON(400, gin.H{
			"error":  err.Error(),
			"status": ok,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success add user",
	})
}

// DeleteUser test
func DeleteUser(c *gin.Context) {
	pathParam := c.Param("id")
	id, _ := strconv.Atoi(pathParam)
	userService := model.Env{DB: db.DB}
	// get id from param
	// id := c.Params.ByName("id")
	rowsAffected, err2 := userService.DeleteUser(id)
	if err2 != nil {
		panic(err2.Error())
	}
	fmt.Println(rowsAffected)
	c.JSON(http.StatusOK, gin.H{
		"message": "delete OK",
	})
}
