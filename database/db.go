package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func CreateDB() {
	db, err := sql.Open("sqlite3", "ToDoDB.sqlite")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	createTable(db)
}

func createTable(db *sql.DB) {
	sqlTable := `
	CREATE TABLE IF NOT EXISTS items(
			ID INTEGER PRIMARY KEY,
			Name TEXT NOT NULL,
			Completed INTEGER NOT NULL
	);
	`

	_, err := db.Exec(sqlTable)

	if err != nil {
		log.Fatal(err)
	}
}

func AddTask(task Task) {
	sqlTask := `
	INSERT OR REPLACE INTO items(
		Name,
		Completed
	) values(?, ?)
	`
	db, _ := sql.Open("sqlite3", "ToDoDB.sqlite")

	statement, err := db.Prepare(sqlTask)

	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()
	defer db.Close()

	_, err = statement.Exec(task.Name, task.Completed)

	if err != nil {
		log.Fatal(err)
	}
}

func deleteTask(db *sql.DB, id int) {
	sqlTask := "DELETE FROM items WHERE id=" + strconv.Itoa(id)

	statement, err := db.Prepare(sqlTask)

	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	_, err = statement.Exec()

	if err != nil {
		log.Fatal(err)
	}

}

func DisplayTasks() []byte {
	db, _ := sql.Open("sqlite3", "ToDoDB.sqlite")

	row, err := db.Query("SELECT * FROM items ORDER BY id")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	defer row.Close()

	var tasks []Task

	for row.Next() {
		var ID int
		var name string
		var completed bool
		row.Scan(&ID, &name, &completed)
		tasks = append(tasks, Task{ID, name, completed})
	}

	response, _ := json.Marshal(tasks)

	fmt.Println(string(response))
	return response

}
