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

	http.HandleFunc("/ws", serveWebsocket)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func checkDatabaseExistence() {
	if _, err := os.Stat("./ToDoDB.sqlite"); err != nil {
		database.CreateDB()
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func serveWebsocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatalln(err)
	}

	defer ws.Close()

	fmt.Println("Connected!")

	reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Println("Error: ", err)
			return
		}

		fmt.Println(messageType, string(p))
		writer(conn)
	}

}

func writer(conn *websocket.Conn) {
	tasks := database.DisplayTasks()
	conn.WriteMessage(websocket.TextMessage, tasks)
}
