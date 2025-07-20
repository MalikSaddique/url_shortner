package api

import (
	"github.com/MalikSaddique/url_shortner/db"
	"github.com/MalikSaddique/url_shortner/models"
)

type URLShortnerAPI interface {
	CreateURL(url models.URL) (*models.URL, error)
}

type URLShortnerAPIImpl struct {
	urldb db.URLDB
}

func NewURLShortnerAPI(input NewURLShortnerAPIInput) URLShortnerAPI {
	return &URLShortnerAPIImpl{
		urldb: input.URLdb,
	}
}

type NewURLShortnerAPIInput struct {
	URLdb db.URLDB
}

var _ URLShortnerAPI = &URLShortnerAPIImpl{}
