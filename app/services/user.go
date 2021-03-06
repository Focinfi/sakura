package services

import (
	"github.com/Focinfi/sakura/app/errors"
	"github.com/Focinfi/sakura/app/i18n"
	"github.com/Focinfi/sakura/app/log"
	"github.com/Focinfi/sakura/app/models"
	"github.com/Focinfi/sakura/app/workers"
	"github.com/Focinfi/sakura/db"
	"github.com/Focinfi/sakura/libs/token"
	"github.com/Focinfi/sakura/libs/utils"
)

// CreateUser creates a new user
func CreateUser(params *models.RequestParams) errors.Error {
	translator := i18n.NewTranslator(i18n.Locale(params.Locale))

	// check params completeness
	if params.User.Password == "" {
		return errors.New(errors.PasswordIsEmpty, "password_is_empty")
	}

	switch params.RegistrationType {
	case models.EmailRegistration:
		if !utils.IsEmail(params.User.Email) {
			return errors.New(errors.EmailIsWrong, "email_is_wrong")
		}
		params.User.Phone = ""
	case models.PhoneRegistration:
		if !utils.IsPhone(params.User.Phone) {
			return errors.New(errors.PhoneIsWrong, "phone_is_wrong")
		}
		// check verification code
		if ok, err := VerifyPhoneCode(params.User.Phone, params.VerificationCode); err != nil {
			return errors.InternalServerError
		} else if !ok {
			return errors.New(errors.PhoneVerificationCodeIsWrong, "phone_verification_code_is_wrong")
		}
		params.User.Email = ""
	}

	// check uniqueness
	if unique, err := params.User.CheckUniqueness(); err != nil {
		return errors.InternalServerError
	} else if !unique {
		return errors.New(errors.UserAlreadyExists, "user_already_exist")
	}

	// create user
	query := db.DB.Create(params.User)
	if query.Error != nil {
		log.DBError(query.Value, query.Error, "failed to create a user")
		return errors.InternalServerError
	}

	// send verification code for email registration
	if params.RegistrationType == models.EmailRegistration {
		tkn, err := token.New("email_verification", 3600*24*30).Set("email", params.User.Email).Sign()
		if err != nil {
			log.LibError("token", "faild to create a email_verification token")
			return errors.InternalServerError
		}

		email := workers.SendEmail(
			translator.T("please_verify_your_email"),
			translator.T("verification_email_boday", tkn),
			params.User.Email,
		)
		workers.EmailWorker.PerformAsync(email)
	}

	return nil
}
