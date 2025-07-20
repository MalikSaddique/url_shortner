package main

import (
	"log"

	"github.com/MalikSaddique/url_shortner/config"
	"github.com/MalikSaddique/url_shortner/db"
	"github.com/MalikSaddique/url_shortner/router"
)

func main() {
	config.LoadConfig()

	conn, err := db.Connection()
	if err != nil {
		log.Fatalf("PostgreSQL connection error: %s", err)
	}
	httpRouter := router.NewRouter(conn)
	if err := httpRouter.Engine.Run(":8003"); err != nil {
		log.Fatalf("HTTP server failed to start: %s", err)
	}

}
