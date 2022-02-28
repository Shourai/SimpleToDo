package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Shourai/SimpleToDo/database"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}

	return conn, nil
}

func Reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Println("Error: ", err)
			return
		}

		fmt.Println(messageType, string(p))
		task := database.Task{
			Name:      string(p),
			Completed: false,
		}
		database.AddTask(task)

		taskJson, _ := json.Marshal(task)

		conn.WriteMessage(websocket.TextMessage, taskJson)

		// writer(conn)
	}

}

func Writer(conn *websocket.Conn) {
	tasks := database.DisplayTasks()
	conn.WriteMessage(websocket.TextMessage, tasks)
}
