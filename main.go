package main

import (
	"net/http"

	"notaku/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()

	// Register middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// session middleware
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// cors middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3500"}
	r.Use(cors.New(config))

	// route middleware
	r.NoRoute(notFound)

	// Register routes
	routes.Routes(r)
	r.NoRoute(notFound)

	r.Run()
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json: "password"`
}

type JWT struct {
	Token string `json: "token"`
}

func Signup(c *gin.Context) {
	// signup
	// handle email and password
	// firstname lastname
	// signup
	// auto login
	session := sessions.Default(c)
	if session.Get("hello") != "world" {
		session.Set("hello", "world")
		session.Save()
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "route exixst",
		"hello":   session.Get("hello"),
	})
}

func Login(c *gin.Context) {
	// accept email
	// accept password
	// generate token
	// pass token to client
	c.JSON(http.StatusOK, gin.H{
		"message": "login exixst",
	})
}

// not found page
func notFound(c *gin.Context) {
	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}

// nomethods commentj
func NoMethods(route *gin.Engine) {
	route.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(405, "not allowed")
	})
}

type QueryError struct {
	Query string
	Err   error
}
