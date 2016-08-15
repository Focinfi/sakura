package services

import "github.com/Focinfi/sakura/app/log"

// PhoneVerifier contains SendCode and VerifyCode
type PhoneVerifier interface {
	SendCode(phone string) error
	VerifyCode(phone, code string) (bool, error)
}

var phoneVerifier = NEPhoneVerifier{}

// NEPhoneVerifier implements PhoneVerifier interface
type NEPhoneVerifier struct {
}

// SendCode sends code for the phone
func (verifier *NEPhoneVerifier) SendCode(phone string) error {
	// TODO: use specific third-party service
	return nil
}

// VerifyCode verify the code for the phone
func (verifier *NEPhoneVerifier) VerifyCode(phone, code string) (bool, error) {
	// TODO: use specific third-party service
	return true, nil
}

// SendPhoneCode sends code
func SendPhoneCode(phone string) error {
	return phoneVerifier.SendCode(phone)
}

// VerifyPhoneCode verifies code
func VerifyPhoneCode(phone, code string) (ok bool, err error) {
	ok, err = phoneVerifier.VerifyCode(phone, code)
	if err != nil {
		log.ThirdPartyServiceError("verify_phone_code", err, nil, phone, code)
	}
	return
}
