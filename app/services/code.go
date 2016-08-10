package services

// Code for response code
type Code int

const (
	// AccessTokenIsWrong for access_token error
	AccessTokenIsWrong = 9000
	// JSONBodyParsingError for json parsing error
	JSONBodyParsingError = 9001
	// ActionIsNotAllowed for action error
	ActionIsNotAllowed = 9002
	// LoginTokenIsWorng for bad login_token
	LoginTokenIsWorng = 9003
	// EmailIsWrong for empty email
	EmailIsWrong = 9004
	// PhoneIsWrong for wrong phone
	PhoneIsWrong = 9005
)
