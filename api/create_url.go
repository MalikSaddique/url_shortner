package api

import (
	"math/rand"

	"github.com/MalikSaddique/url_shortner/models"
)

func (api *URLShortnerAPIImpl) CreateURL(url models.URL) (*models.URL, error) {
	if url.Long_URL[:4] != "http" && url.Long_URL[:5] != "https" {
		url.Long_URL = "http://" + url.Long_URL
	}
	key := generateKey(8)

	response, err := api.urldb.CreateURL(url, key)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func generateKey(length int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
