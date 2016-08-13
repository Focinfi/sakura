package services

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

// SendCode sends code
func SendCode(phone string) error {
	return phoneVerifier.SendCode(phone)
}

// VerifyCode verifies code
func VerifyCode(phone, code string) (bool, error) {
	return phoneVerifier.VerifyCode(phone, code)
}
