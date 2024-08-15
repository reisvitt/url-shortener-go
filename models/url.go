package models

import (
	"time"
)

type Url struct {
	ID        string    `bson:"_id" json:"id"`
	FullUrl   string    `bson:"name" json:"full-url"`
	ExpiresAt time.Time `bson:"expires-at" json:"expires-at"`
	CreatedAt time.Time `bson:"created-at" json:"created-at"`
}
