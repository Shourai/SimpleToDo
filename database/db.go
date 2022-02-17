package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type task struct {
	Name string
	Done int
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

	addTask(db, task{Name: "First task", Done: 0})
}

func createTable(db *sql.DB) {
	sqlTable := `
	CREATE TABLE IF NOT EXISTS items(
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
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
