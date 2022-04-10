package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func StartServer(mux *mux.Router, address, cert, key string) (*http.Server, chan error) {
	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}
	shutdown := make(chan error)
	go func() {
		if cert != "" || key != "" {
			err := server.ListenAndServeTLS(cert, key)
			shutdown <- err
		} else {
			err := server.ListenAndServe()
			shutdown <- err
		}
	}()
	return server, shutdown
}

func waitForServerToClose(shutdown <-chan error) error {
	err := <-shutdown
	if err == http.ErrServerClosed {
		return nil
	}
	return err
}
