package http

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func RunServer(router http.Handler) {
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Running server on '%s'\n", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}
