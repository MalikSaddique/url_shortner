package models

import "time"

type URL struct {
	Id         int64     `json:"id"`
	Short_URL  string    `json:"shorturl"`
	Long_URL   string    `json:"longurl"`
	Created_at time.Time `json:"time"`
}
