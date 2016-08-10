package services

import (
	"github.com/Focinfi/sakura/app/errors"
	"github.com/Focinfi/sakura/app/models"
	"github.com/Focinfi/sakura/libs/utils"
)

// CreateUser creates a new user
func CreateUser(params *models.RequestParams) errors.Error {

	switch params.RegistrationType {
	case models.EmailRegistration:
		if !utils.IsEmail(params.User.Email) {
			return errors.New(errors.EmailIsWrong, "email_is_wrong")
		}
	case models.PhoneRegistration:
		if !utils.IsPhone(params.User.Phone) {
			return errors.New(errors.PhoneIsWrong, "phone_is_wrong")
		}
		// check verification code
		if ok, err := phoneVerifier.VerifyCode(params.User.Phone, params.VerificationCode); err != nil {
			return errors.InternalServerError
		} else if !ok {
			return errors.New(errors.PhoneVerificationCodeIsWrong, "phone_verification_code_is_wrong")
		}
	}

	return nil
}
