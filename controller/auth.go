package controller

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){
	fmt.Println("success")
	c.JSONP(http.StatusOK, gin.H{
		"message": "login",
	})
}

func Signup(c *gin.Context){
	fmt.Println("success")
	c.JSONP(http.StatusOK, gin.H{
		"message": "signup",
	})
}


func Logout(c *gin.Context){
	fmt.Println("success")
	c.JSONP(http.StatusOK, gin.H{
		"message": "logout",
	})
}