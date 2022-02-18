package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type task struct {
	ID   int
	Name string
	Done bool
}

func CreateDB() {
	db, err := sql.Open("sqlite3", "ToDoDB.sqlite")

	// var ctx context.Context = context.Background()
	// ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	// defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	createTable(db)

	// addTask(db, task{Name: "First task", Done: false})
	// deleteTask(db, 2)
	displayTasks(db)

}

func createTable(db *sql.DB) {
	sqlTable := `
	CREATE TABLE IF NOT EXISTS items(
			ID INTEGER PRIMARY KEY,
			Name TEXT NOT NULL,
			Done INTEGER NOT NULL
	);
	`

	_, err := db.Exec(sqlTable)

	if err != nil {
		log.Fatal(err)
	}
}

func addTask(db *sql.DB, task task) {
	sqlTask := `
	INSERT OR REPLACE INTO items(
		Name,
		Done
	) values(?, ?)
	`

	statement, err := db.Prepare(sqlTask)

	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	_, err = statement.Exec(task.Name, task.Done)

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

func displayTasks(db *sql.DB) {
	row, err := db.Query("SELECT * FROM items ORDER BY id")

	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	var tasks []task

	for row.Next() {
		var ID int
		var name string
		var done bool
		row.Scan(&ID, &name, &done)
		tasks = append(tasks, task{ID, name, done})
	}

	fmt.Println(tasks)

}
