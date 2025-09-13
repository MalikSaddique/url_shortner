package models

import "time"

type URL struct {
	Id         int64     `json:"id"`
	ShortURL   string    `json:"shorturl"`
	LongURL    string    `json:"longurl"`
	ExpireAt   time.Time `json:"expire_at"`
	VisitCount int64     `json:"visit_count"`
	CreatedAt  time.Time `json:"time"`
}
type CreateURLResponse struct {
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}
type CreateURLRequest struct {
	LongURL string `json:"long_url"`
}
