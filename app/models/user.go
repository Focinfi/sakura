package models

import (
	"github.com/jinzhu/gorm"
)

// User for users
type User struct {
	gorm.Model
	UUID     string `json:"uuid" gorm:"type:char(32)"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"type:varchar(100);unique_index"`
	Phone    string `json:"phone" gorm:"type:varchar(20);unique_index"`
	Verified bool   `json:"verified"`
}

// DisplayName implements Auth interface
func (user *User) DisplayName() string {
	return user.Name
}
