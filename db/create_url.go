package db

import (
	"log"

	"github.com/MalikSaddique/url_shortner/models"
)

func (db *URLDBImpl) CreateURL(url models.URL, key string) (*models.URL, error) {

	_, err := db.Db.Exec("INSERT INTO url (longurl) VALUES ($1)", url.Long_URL)
	if err != nil {
		log.Println("Failed to create URL")
		return nil, err
	}

	return &url, nil
}
