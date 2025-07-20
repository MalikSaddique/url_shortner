package api

import (
	"github.com/MalikSaddique/url_shortner/db"
)

type URLShortnerAPI interface {
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
