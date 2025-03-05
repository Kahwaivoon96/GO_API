package database

import (
	"database/sql"
	config "example/web-service-gin/config"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB // Global database connection

// Connect initializes the MSSQL connection once
func Connect(cfg *config.Config) error {

	dsn := fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s;encrypt=disable",
		cfg.Database.Server, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)

	var err error
	DB, err = sql.Open("sqlserver", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("database is not responding: %w", err)
	}

	log.Println("✅ Connected to MSSQL database successfully")
	return nil
}
