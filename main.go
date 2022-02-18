package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Shourai/SimpleToDo/database"
	"github.com/gorilla/websocket"
)

func main() {
	checkDatabaseExistence()

	database.DisplayTasks()

	http.HandleFunc("/ws", serveWebsocket)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func checkDatabaseExistence() {
	if _, err := os.Stat("./ToDoDB.sqlite"); err != nil {
		database.CreateDB()
	}
}

var upgrader = websocket.Upgrader{}

func serveWebsocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatalln(err)
	}

	defer ws.Close()

	fmt.Println("Connected!")

	readIncoming(ws)
}

func readIncoming(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(messageType, string(p))
	}

}
