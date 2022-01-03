package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Start(route *mux.Router, endpoint string) {
	server := &http.Server{
		Handler: route,
		Addr:    endpoint,
	}

	fmt.Println("Server running on ", endpoint)

	server.ListenAndServe()
}
