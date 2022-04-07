package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RegisterRequest struct {
	Username string `json:"username"`

	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`

	Password string `json:"password"`
}

type LoginResponse struct {
}

func (a *API) handleRegister(w http.ResponseWriter, r *http.Request) {

	r, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.errorResponse(w, r.URL.Path, http.StatusInternalServerError, "", err)
		return
	}

	var login LoginRequest
	err = json.Unmarshal(r.Body, &login)

	if err != nil {

	}

}

func (a *API) handleLogin(w http.ResponseWriter, r *http.Request) {

}
