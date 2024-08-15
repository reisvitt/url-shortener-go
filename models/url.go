package models

import "time"

type Url struct {
	ID        string    `json:"id"`
	FullUrl   string    `json:"full-url"`
	ExpiresAt time.Time `json:"expires-at"`
}
