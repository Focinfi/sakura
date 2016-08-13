package handlers

import (
	"github.com/Focinfi/sakura/app/errors"
	"github.com/Focinfi/sakura/app/i18n"
	"github.com/Focinfi/sakura/app/models"
	"github.com/Focinfi/sakura/app/response"
	"github.com/Focinfi/sakura/app/services"
	"github.com/Focinfi/sakura/libs/utils"
)

// SendPhoneVerificationCode for send phone verification code action
func SendPhoneVerificationCode(c *models.Context) {
	if !utils.IsPhone(c.Params.Phone) {
		response.Failed(c.Context,
			errors.PhoneIsWrong,
			i18n.T(c.Params.Locale, "phone_is_wrong"))
		return
	}

	if err := services.SendCode(c.Params.VerificationCode); err != nil {
		response.Failed(c.Context,
			errors.InternalServerError.Code(),
			errors.InternalServerError.Message(c.Params.Locale))
		return
	}

	response.OK(c.Context, nil)
}
