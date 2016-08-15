package models

import (
	"github.com/Focinfi/sakura/app/log"
	"github.com/Focinfi/sakura/db"
	"github.com/Focinfi/sakura/libs/utils"
	"github.com/icrowley/fake"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// User for users
type User struct {
	gorm.Model
	UUID     string `json:"uuid" gorm:"type:char(32)"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"type:varchar(100);unique_index"`
	Phone    string `json:"phone" gorm:"type:varchar(20);unique_index"`
	Password string `json:"password"`
	Solt     string `json:"solt"`
	Verified bool   `json:"verified"`
}

// DisplayName implements Auth interface
func (user *User) DisplayName() string {
	return user.Name
}

// CheckUniqueness for new User
func (user *User) CheckUniqueness() (bool, error) {
	u := &User{}
	query := db.DB.Where(&User{Email: user.Email, Phone: user.Phone}).First(u)
	if query.Error != nil {
		log.DBError(query.Value, query.Error, "failed to get user")
	}
	return u.ID == 0, query.Error
}

// BeforeCreate adds uuid and set default name for new user
func (user *User) BeforeCreate(database *gorm.DB) error {
	if user.UUID != "" {
		user.UUID = uuid.NewV4().String()
	}

	if user.Name == "" {
		user.Name = fake.FullName()
	}

	if user.Password != "" {
		user.Solt = utils.RandCharString(16)
		user.Password = utils.Sha256(user.Password + user.Solt)
	}

	return nil
}
