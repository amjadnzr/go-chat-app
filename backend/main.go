package main

import (
	"fmt"
	"net/http"

	"github.com/amjadnzr/go-chat-app/websocket"
	"githun.com/amjadnzr/go-chat-app/pkg/websocket"
)

func main() {
	fmt.Println("Amjad's Full Stack Project")
	setUpRoutes()
	http.ListenAndServe(":9000", nil)

}

func setUpRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Web Socket endpoint reached")
	conn, err := websocket.Upgrage(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()

}
