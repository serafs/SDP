package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type API struct {
	logger *log.Logger
}

func NewAPI() *API {
	return &API{}
}

func (a *API) RegisterRoutes(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()
	api.Use()

	api.HandleFunc("/register", a.handleRegister).Methods("POST")
	api.HandleFunc("/login", a.handleLogin).Methods("POST")
}

func (a *API) errorResponse(w http.ResponseWriter, api string, code int, message string, sourceError error) {
	a.logger.Error("API ERROR",
		a.logger.Int("code", code),
		a.logger.Err(sourceError),
		a.logger.String("msg", message),
		a.mlog.String("api", api),
	)
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(model.ErrorResponse{Error: message, ErrorCode: code})
	if err != nil {
		data = []byte("{}")
	}
	w.WriteHeader(code)
	_, _ = w.Write(data)
}
