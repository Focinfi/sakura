package errors

import "github.com/Focinfi/sakura/app/i18n"

// Error represents errors of services
type Error interface {
	Code() int
	Message(locale string) string
}

type err struct {
	code       Code
	messageKey string
}

func (e err) Code() int {
	return int(e.code)
}

func (e err) Message(locale string) string {
	return i18n.T(i18n.Locale(locale), e.messageKey)
}

// New allocates and returns a new ErrorFunc
func New(code Code, messageKey string) Error {
	return err{code: code, messageKey: messageKey}
}

// InternalServerError for internal server error
var InternalServerError = New(Code(500), "internal_server_error_please_retry")
