package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"homework-4/internal/config"
	"homework-4/internal/server"
	"log"
	"os"
	"time"

	// Import the pq driver so that it can register itself with the database/sql
	// package. Note that we alias this import to the blank identifier, to stop the Go
	//compiler complaining that the package isn't being used.
	_ "github.com/lib/pq"
)

type application struct {
	config config.Config
	logger *log.Logger
}

func main() {
	cfg := config.NewConfig()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	fmt.Printf("%+v", cfg)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}(db)

	e := echo.New()
	s := server.NewUserServer(db)
	server.RegisterHandlers(e, s)

	server.RegisterHealthHandler(e)
	server.RegisterHostnameHandler(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
}

func openDB(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
