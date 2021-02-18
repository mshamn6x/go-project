package main

import (
	"log"
	"net/http"
	"new/test/project/ui/router"
	"time"
)

func main() {

	r := router.New()

	// This will serve files under http://localhost:8000/static/<filename>
	r.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir(`C:\Users\mshanm6x\Downloads\project\ui\public`))))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
