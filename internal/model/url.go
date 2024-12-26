package model

import "time"

type URL struct {
	ID string `json:"id" bson:"id"`
	LongURL string `json:"longUrl" bson:"longUrl"`
	ShortURL string `json:"shortUrl" bson:"shortUrl"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt" bson:"expiresAt"`
}