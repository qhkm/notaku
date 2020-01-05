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

// PostService test
type PostService struct{}

// Env test
type Env struct {
	db *sql.DB
}

// GetPostList test
func (p *PostService) GetPostList() []Post {
	db, err := sql.Open("sqlite3", "./note.db")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()
	results, err := db.Query("SELECT * FROM posts")
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
