package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type task struct {
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

	task1 := task{Name: "Task one", Completed: false}
	task2 := task{Name: "Task two", Completed: false}
	task3 := task{Name: "Task three", Completed: false}
	task4 := task{Name: "Task four", Completed: true}
	task5 := task{Name: "Task five", Completed: false}

	addTask(db, task1)
	addTask(db, task2)
	addTask(db, task3)
	addTask(db, task4)
	addTask(db, task5)

}

func addTask(db *sql.DB, task task) {
	sqlTask := `
	INSERT OR REPLACE INTO items(
		Name,
		Completed
	) values(?, ?)
	`

	statement, err := db.Prepare(sqlTask)

	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

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

	var tasks []task

	for row.Next() {
		var ID int
		var name string
		var completed bool
		row.Scan(&ID, &name, &completed)
		tasks = append(tasks, task{ID, name, completed})
	}

	response, _ := json.Marshal(tasks)

	fmt.Println(string(response))
	return response

}
