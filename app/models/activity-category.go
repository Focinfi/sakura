package models

import "github.com/jinzhu/gorm"

// ActivityCategory for activity_category
type ActivityCategory struct {
	gorm.Model
	Name string `json:"name"`
}
