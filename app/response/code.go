package response

// Code for response code
type Code int

const (
	// StatusOK for 200
	StatusOK Code = 200
	// StatusInternalServerError for 500
	StatusInternalServerError = 500
	// AccessTokenIsWrong for access_token error
	AccessTokenIsWrong = 9000
	// JSONBodyParsingError for json parsing error
	JSONBodyParsingError = 9001
	// ActionIsNotAllowed for action error
	ActionIsNotAllowed = 9002
)
