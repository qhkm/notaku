package model

import (
	"database/sql"
	"fmt"
)

// Post int
type Post struct {
	ID    int
	Title string
	Body  string
}

// Env test
type Env struct {
	DB *sql.DB
}

// GetPostList test
func (p *Env) GetPostList() []Post {
	results, err := p.DB.Query("SELECT * FROM posts")
	if err != nil {
		panic(err.Error())
	}

	// var posts []model.Post
	var posts []Post
	for results.Next() {
		var post Post
		err := results.Scan(&post.ID, &post.Title, &post.Body)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
		// fmt.Println(strconv.Itoa(post.Id) + ": " + post.Title + " " + post.Body)
	}
	fmt.Println(posts)
	return posts

}

// type User struct {
// 	Id   int
// 	Name string
// 	Age  int
// }
