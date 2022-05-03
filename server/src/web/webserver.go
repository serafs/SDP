package web

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

type RoutedService interface {
	RegisterRoutes(router *mux.Router)
}

type Server struct {
	http.Server
	baseUrl  string
	port     int
	rootPath string
}

func (ws *Server) Router() *mux.Router {
	return ws.Server.Handler.(*mux.Router)
}

func (ws *Server) registerRoutes() {
	ws.Router().PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(ws.rootPath, "static")))))
	ws.Router().PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		indexTemplate, err := template.New("index").ParseFiles(path.Join(ws.rootPath, "index.html"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = indexTemplate.ExecuteTemplate(w, "index.html", map[string]string{"BaseURL": ws.baseUrl})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func (ws *Server) Shutdown() error {
	return ws.Close()
}

// fileExists returns true if a file exists at the path.
func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return err == nil
}
