package main

import (

	_ "github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"notaku/routes"
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
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3500"}
	// r.Use(cors.New(config))

	// route middleware
	r.NoRoute(notFound)

	// Register routes
	routes.Routes(r)
	r.NoRoute(notFound)

	r.Run()
}

// not found page
func notFound(c *gin.Context) {
	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}

// NoMethods test
func NoMethods(route *gin.Engine) {
	route.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(405, "not allowed")
	})
}

// QueryError test
type QueryError struct {
	Query string
	Err   error
}
