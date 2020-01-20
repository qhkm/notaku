package model

import "fmt"

type User struct {
	ID   int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int `json:"age,omitempty"`
}


// AllPosts test
func (env *Env) AllUsers() []User {
	results, err := env.DB.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	// var posts []model.Post
	var users []User
	for results.Next() {
		var user User
		err := results.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
		// fmt.Println(strconv.Itoa(post.Id) + ": " + post.Title + " " + post.Body)
	}
	fmt.Println(users)
	return users

}