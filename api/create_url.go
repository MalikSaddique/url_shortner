package api

import (
	"math/rand"
	"strings"
	"time"

	"github.com/MalikSaddique/url_shortner/models"
)

func (api *URLShortnerAPIImpl) CreateURL(url models.URL) (*models.URL, error) {
	if !strings.HasPrefix(url.LongURL, "http://") && !strings.HasPrefix(url.LongURL, "https://") {
		url.LongURL = "http://" + url.LongURL
	}

	key := generateKey(8)

	url.ShortURL = key
	url.CreatedAt = time.Now()
	if url.ExpireAt.IsZero() {
		url.ExpireAt = time.Now().Add(30 * 24 * time.Hour)
	}

	response, err := api.urldb.CreateURL(url)
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
