package ws

import (
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

func (ws *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {

}

func (ws *Server) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/ws", ws.handleWebSocket)
}

func (ws *Server) addListener(client *Client) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.listeners[client] = true
}
