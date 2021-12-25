package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/serafs/SDP/server/core-service/ws"
	"io/ioutil"
	"net/http"
)

type msg struct {
	Num int
}

func main() {

	wsServer := ws.NewServer()
	http.HandleFunc("/", rootHandler)

	panic(http.ListenAndServe(":9000", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}

//func wsHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Header.Get("Origin") != "http://"+r.Host {
//		http.Error(w, "Origin not allowed", 403)
//		return
//	}
//	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
//	if err != nil {
//		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
//	}
//
//	ws.NewServer()
//
//	go echo(conn)
//}

func echo(conn *websocket.Conn) {
	for {
		m := msg{}

		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}

		fmt.Printf("Got message: %#v\n", m)

		if err = conn.WriteJSON(m); err != nil {
			fmt.Println(err)
		}
	}
}
