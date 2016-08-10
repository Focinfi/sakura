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
	return nil
}

// VerifyCode verify the code for the phone
func (verifier *NEPhoneVerifier) VerifyCode(phone, code string) (bool, error) {
	return true, nil
}
