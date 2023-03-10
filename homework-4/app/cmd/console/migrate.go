package main

import (
	"fmt"
	"log"
	"os"

	"homework-4/internal/config"

	_ "github.com/lib/pq"
	goose "github.com/pressly/goose/v3"
)

func main() {
	cfg := config.NewConfig()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	fmt.Printf("%+v", cfg)

	db, err := goose.OpenDBWithDriver("postgres", cfg.DB.DSN)
	if err != nil {
		logger.Fatal(err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		logger.Fatal(err)
	}

	if err := goose.Up(db, cfg.DB.MigrationsSource); err != nil {
		logger.Fatal(err)
	}
}
