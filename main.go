package main

import (
	"fmt"
	"net/http"

	"github.com/MalikSaddique/url_shortner/config"
	"github.com/MalikSaddique/url_shortner/db"
)

func main() {
	config.LoadConfig()

	// conn, err := db.Connection()
	// if err != nil {
	// 	log.Fatalf("PostgreSQL connection error: %s", err)
	// }
	db.Connection()

	// if err := Engine.Run(":8005"); err != nil {
	// 	log.Fatalf("HTTP server failed to start: %s", err)
	// }
	fmt.Println("Server listening on port 8005")
	http.ListenAndServe(":8005", nil)

}
