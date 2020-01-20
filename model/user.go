package model

import (
	"fmt"
	"log"
)

type User struct {
	ID   int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int `json:"age,omitempty"`
}

// Single user test
func (env *Env) SingleUser(id int) User {
	sqlStatement := "SELECT * FROM users WHERE id = ?"
	results, err := env.DB.Query(sqlStatement, id)
	if err != nil {
		log.Fatal(err)
	}

	var user User
	for results.Next() {
		err := results.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(user)
	return user

}


func (env *Env) AllUsers() []User {
	results, err := env.DB.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	var users []User
	for results.Next() {
		var user User
		err := results.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	fmt.Println(users)
	return users

}


func (env *Env) UpdateUser(user User) (int64, error) {
	sqlStatement := `UPDATE users SET name = ?, age = ? WHERE id = ?`
	fmt.Println(user)
	results, err := env.DB.Exec(sqlStatement, user.Name, user.Age, user.ID)
	if err != nil {
		panic(err.Error())
	}
	rowsAffected, _ := results.RowsAffected()
	return rowsAffected, nil
}

// AddUser test
func (env *Env) AddUser(user User) (int64, error) {
	sqlStatement := `INSERT INTO users(name, age) values(?, ?)`
	results, err := env.DB.Exec(sqlStatement, user.Name, user.Age)
	if err != nil {
		panic(err.Error())
	}
	id, _ := results.LastInsertId()
	rowsAffected, _ := results.RowsAffected()
	fmt.Println(id)
	return rowsAffected, nil
}

// DeleteUser test
func (env *Env) DeleteUser(id int) (int64, error) {

	sqlStatement := `DELETE FROM users where id = ?`
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
