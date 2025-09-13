package db

import (
	"log"
	"time"

	"github.com/MalikSaddique/url_shortner/models"
)

func (db *URLDBImpl) CreateURL(url models.URL) (*models.URL, error) {

	err := db.Db.QueryRow("INSERT INTO url (shorturl,longurl,time) VALUES ($1,$2,$3) RETURNING id", url.ShortURL, url.LongURL, time.Now()).Scan(&url.Id)
	if err != nil {
		log.Println("Failed to create URL")
		return nil, err
	}

	return &url, nil
}
