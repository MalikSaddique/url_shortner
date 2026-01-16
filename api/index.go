package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MalikSaddique/url_shortner/config"
	"github.com/MalikSaddique/url_shortner/db"
	"github.com/MalikSaddique/url_shortner/router"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	config.LoadConfig()

	conn, err := db.Connection()
	if err != nil {
		log.Printf("PostgreSQL connection error: %s", err)
		http.Error(w, fmt.Sprintf("Database connection error: %s", err), http.StatusInternalServerError)
		return
	}

	httpRouter := router.NewRouter(conn)
	httpRouter.Engine.ServeHTTP(w, r)
}
