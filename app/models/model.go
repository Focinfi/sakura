package models

import (
	"time"
)

// Model for base model
type Model struct {
	ID        string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
