package database

import (
	"database/sql"
	"fmt"
	"os"
	"sqap/internal/config"
)

func ConnectDb(cfg *config.Config) *sql.DB {
	db, err := sql.Open("libsql", cfg.DBUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to db %s: %s", cfg.DBUrl, err)
		os.Exit(1)
	}

	return db
}
