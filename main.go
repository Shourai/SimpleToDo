package main

import (
	"log"
	"net/http"

	"github.com/Shourai/SimpleToDo/database"
	websocket "github.com/Shourai/SimpleToDo/websocket"
)

func serveWebsocket(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()

	// conn.WriteMessage(1, database.DisplayTasks())
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWebsocket(pool, w, r)
	})
	// http.HandleFunc("/ws", serveWebsocket)
}

func main() {
	database.CheckDatabaseExistence()
	setupRoutes()

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	// log.Fatal(http.ListenAndServeTLS("localhost:8080"))
}
