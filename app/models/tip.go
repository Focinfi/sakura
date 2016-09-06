package models

import "github.com/jinzhu/gorm"

// Tip for tip model
type Tip struct {
	gorm.Model
	Content string `json:"content"`
}
