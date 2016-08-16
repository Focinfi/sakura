package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ActivityNote for user activity note
type ActivityNote struct {
	gorm.Model
	UserID       uint       `json:"user_id"`
	Time         *time.Time `json:"time"`
	Duration     uint       `json:"duration"`
	Content      string     `json:"content"`
	Satisfaction int        `json:"satisfaction"`
}
