package api

import "github.com/MalikSaddique/url_shortner/models"

func (api *URLShortnerAPIImpl) GetByShortURL(short string) (*models.URL, error) {
	return api.urldb.GetByShortURL(short)
}
