package model

import "time"

type APIKey struct {
	ID        int
	Key       string
	CreatedAt time.Time
}
