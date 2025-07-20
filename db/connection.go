package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/MalikSaddique/url_shortner/config"
)

func Connection() (*sql.DB, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	host := cfg.URLShortner.POSTGRES_HOST
	port := cfg.URLShortner.POSTGRES_PORT
	user := cfg.URLShortner.POSTGRES_USER
	password := cfg.URLShortner.POSTGRES_PASSWORD
	dbname := cfg.URLShortner.POSTGRES_DB

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("Database credentials are not fully set in the environment variables")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB connection: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %v", err)
	}

	log.Println("Postgres is connected. Database is connected.")

	return DB, nil
}
