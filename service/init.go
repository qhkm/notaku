package service

import "database/sql"

func init() {
	//connect to db
	db, err := sql.Open("sqlite3", "./note.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS posts (id INTEGER PRIMARY KEY, title TEXT, body TEXT)")
	// statement.Exec()
	// statement, _ = database.Prepare("INSERT INTO posts (title, body) VALUES (?, ?)")
	// statement.Exec("title test", "title body")
	// rows, _ := database.Query("SELECT * FROM posts")
	// var id int
	// var title string
	// var body string
	// for rows.Next() {
	// 	rows.Scan(&id, &title, &body)
	// 	fmt.Println(strconv.Itoa(id) + ": " + title + " " + body)
	// }
}
