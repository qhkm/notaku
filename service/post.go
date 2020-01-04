package service

import (
	"database/sql"
	"fmt"
	"notaku/model"
)

// PostService test
type PostService struct{}

// GetPostList test
func (p *PostService) GetPostList() []model.Post {
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
	var posts []model.Post
	for results.Next() {
		var post model.Post
		err := results.Scan(&post.Id, &post.Title, &post.Body)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
		// fmt.Println(strconv.Itoa(post.Id) + ": " + post.Title + " " + post.Body)
	}
	fmt.Println(posts)
	return posts

}
