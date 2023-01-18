package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	port   uint16
	logger *log.Logger
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		port:   8000,
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", app.healthcheckHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting server on %s", srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
