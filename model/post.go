package model

import (
	"database/sql"
	"fmt"
	"log"
)

// Post int
type Post struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

// Env test
type Env struct {
	DB *sql.DB
}

// SinglePost test
func (env *Env) SinglePost(id int) Post {
	sqlStatement := "SELECT * FROM posts WHERE id = ?"
	results, err := env.DB.Query(sqlStatement, id)
	if err != nil {
		log.Fatal(err)
	}

	// var posts []model.Post
	var post Post
	for results.Next() {
		err := results.Scan(&post.ID, &post.Title, &post.Body)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(strconv.Itoa(post.Id) + ": " + post.Title + " " + post.Body)
	}
	fmt.Println(post)
	return post

}

// AllPosts test
func (env *Env) AllPosts() []Post {
	results, err := env.DB.Query("SELECT * FROM posts")
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

// UpdatePost test
func (env *Env) UpdatePost(post Post) (int64, error) {
	sqlStatement := `UPDATE posts SET title = ?, body = ? WHERE id = ?`
	fmt.Println(post)
	results, err := env.DB.Exec(sqlStatement, post.Title, post.Body, post.ID)
	if err != nil {
		panic(err.Error())
	}
	rowsAffected, _ := results.RowsAffected()
	return rowsAffected, nil
}

// AddPost test
func (env *Env) AddPost(post Post) (int64, error) {
	sqlStatement := `INSERT INTO posts(title, body) values(?, ?)`
	results, err := env.DB.Exec(sqlStatement, post.Title, post.Body)
	if err != nil {
		panic(err.Error())
	}
	id, _ := results.LastInsertId()
	rowsAffected, _ := results.RowsAffected()
	fmt.Println(id)
	return rowsAffected, nil
}

// DeletePost test
func (env *Env) DeletePost(id int) (int64, error) {

	sqlStatement := `DELETE FROM posts where id = ?`
	results, err := env.DB.Exec(sqlStatement, id)
	if err != nil {
		panic(err.Error())
	}

	count, err := results.RowsAffected()
	if err != nil {
		return 0, err
	}
	fmt.Println(count)
	return count, nil
}

// type User struct {
// 	Id   int
// 	Name string
// 	Age  int
// }
