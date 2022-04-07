package api

import "github.com/gorilla/mux"

type API struct{}

func (a *API) RegisterRoutes(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()
	api.Use()

	api.HandleFunc("/register", a.handleRegister).Methods("POST")
	api.HandleFunc("/login", a.handleLogin).Methods("POST")
}
