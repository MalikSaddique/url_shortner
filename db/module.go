package db

import (
	"database/sql"

	"github.com/MalikSaddique/url_shortner/models"
)

type URLDB interface {
	CreateURL(url models.URL) (*models.URL, error)
}

type URLDBImpl struct {
	Db *sql.DB
}

func NewURLDB(db *sql.DB) URLDB {
	return &URLDBImpl{
		Db: db,
	}
}
