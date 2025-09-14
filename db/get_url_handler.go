package db

import (
	"database/sql"

	"github.com/MalikSaddique/url_shortner/models"
)

func (db *URLDBImpl) GetByShortURL(short string) (*models.URL, error) {
	var url models.URL

	query := "SELECT id, shorturl, longurl, time FROM url WHERE shorturl = $1 LIMIT 1"
	err := db.Db.QueryRow(query, short).Scan(
		&url.Id,
		&url.ShortURL,
		&url.LongURL,
		&url.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &url, nil
}
