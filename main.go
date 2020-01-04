package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"notaku/controller"
)

func main() {

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api := r.Group("api")
	{
		api.GET("/post/id", controller.GetPost)
		api.GET("/posts", controller.GetAllPosts)
		api.GET("/posts/add", controller.AddPost)
		api.GET("/posts/delete", controller.DeletePost)
	}
	r.Run()
}

//check error
func checkError(err error) {
	if err != nil {
		panic(err)
	}

}

// func handleGet(c *gin.Context) {
// 	listOfPost := []Post{
// 		Post{
// 			Id:    1,
// 			Title: "name",
// 			Body:  "body",
// 		},
// 		Post{
// 			Id:    2,
// 			Title: "name2",
// 			Body:  "body2",
// 		},
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": listOfPost,
// 	})
// }

// func handlePost(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "handlePost",
// 	})
// }
// func handlePut(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "handlePut",
// 	})
// }
// func handleDelete(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "handleDelete",
// 	})
// }

// func getPong(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "pong",
// 	})
// }

// func getTest(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "test",
// 	})
// }

// not found page
func NotFound(route *gin.Engine) {
	route.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(404, "Not Found")
	})
}

// nomethods commentj
func NoMethods(route *gin.Engine) {
	route.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(405, "not allowed")
	})
}

// write code to open file

type QueryError struct {
	Query string
	Err   error
}

// how to declare multiple var in one liner
// var a int
// var a,b int
// var a int, b int
// var (
// 	x int
// 	y int
// )

// func Login(email, password string) (map[string]interface{}){

// 	account := &Account{}
// 	err := GetDB().Table("accounts").Where("email = ? ", email).First(account).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return u.
// 		}
// 	}
// }

// its not very hard. Just declare the var
