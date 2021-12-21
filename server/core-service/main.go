package main

import (
	"fmt"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "main page")
}

func setupRoutes() {
	http.HandleFunc("/", mainPage)
}

func main() {
	fmt.Println("Golang app started on port 9000")
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}
