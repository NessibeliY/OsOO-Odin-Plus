package server

import (
	"fmt"
	"net/http"
)

func RunServer(mux http.Handler) error {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Printf("starting server at http://localhost%s\n", srv.Addr)
	err := srv.ListenAndServe()
	return err
}
