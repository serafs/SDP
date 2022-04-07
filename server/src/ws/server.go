package ws

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type Client struct {
	*websocket.Conn
	mu sync.Mutex
}

type Server struct {
	http.Server
	upgrader  websocket.Upgrader
	mu        sync.RWMutex
	listeners map[*Client]bool
}

func NewServer() *Server {
	return &Server{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		listeners: make(map[*Client]bool),
	}
}

func (ws *Server) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/ws", ws.handleWebSocket)
	r.Handle("/", http.FileServer(http.Dir("templates/")))
}

func (ws *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Errorf("error upgrading to websocket", err)
	}
	client := Client{conn, sync.Mutex{}}

	for {
		_, p, err := client.ReadMessage()
		if err != nil {
			fmt.Errorf("ERROR WebSocket", "client", client.RemoteAddr(), err)
			break
		}
		fmt.Println(p)
	}

	ws.addListener(&client)

}

func (ws *Server) addListener(client *Client) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.listeners[client] = true
}

func (ws *Server) deleteListener(client *Client) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	delete(ws.listeners, client)
}
