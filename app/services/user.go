package services

import (
	"github.com/Focinfi/sakura/app/models"
	"github.com/Focinfi/sakura/libs/utils"
)

// CreateUser creates a new user
func CreateUser(registrationType int, user *models.User) Error {
	switch registrationType {
	case models.EmailRegistration:
		if !utils.IsEmail(user.Email) {
			return NewError(EmailIsWrong, "email_is_wrong")
		}
	case models.PhoneRegistration:
		if !utils.IsPhone(user.Phone) {
			return NewError(PhoneIsWrong, "phone_is_wrong")
		}
	}

	return nil
}
