package main

//var upgrader = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//}

// three types of servers :
// 1. Core API server
// 2. WebServer(static,main page)
// 3. WebSocket Server
func main() {

}

//http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
//	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
//
//	for {
//		// Read message from browser
//		msgType, msg, err := conn.ReadMessage()
//		if err != nil {
//			return
//		}
//
//		// Print the message to the console
//		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
//
//		// Write message back to browser
//		if err = conn.WriteMessage(msgType, msg); err != nil {
//			return
//		}
//	}
//})
//
//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("Here")
//	http.ServeFile(w, r, "./templates/index.html")
//})
//
//http.ListenAndServe(":8080", nil)
