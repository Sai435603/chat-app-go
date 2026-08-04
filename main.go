// followed yt tot #Kcode
package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

//TODO
//[X] HTTP Server
//[X] Upgrade it to websockets once client connects
//[] Add newly connected ws to server
//[] Add ws client
//[] remove client on disconnect
//send broadcasing msg -> no race conditions

var (
	WSPort string = ":8080"
)

type Client struct {
	ID   uuid.UUID
	mu   *sync.RWMutex
	conn *websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		ID:   uuid.New(),
		mu:   new(sync.RWMutex),
		conn: conn,
	}
}

type Server struct {
	clients []*Client
	mu      *sync.RWMutex
}

func NewServer() *Server {
	return &Server{
		clients: []*Client{},
		mu:      new(sync.RWMutex),
	}
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  512,
		WriteBufferSize: 512,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func main() {
	http.HandleFunc("/", handleWS)
	log.Fatal(http.ListenAndServe(WSPort, nil))
}
