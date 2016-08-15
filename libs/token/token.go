package token

import (
	"fmt"
	"time"

	"github.com/Focinfi/sakura/config"
	"github.com/Focinfi/sakura/libs/utils"
	"github.com/dgrijalva/jwt-go"
)

// Token composing *jwt.Token for customization
type Token struct {
	*jwt.Token
}

// New alloates and returns a new Token with name and expiration mins
func New(name string, expMins int) *Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"exp":  time.Now().Add(time.Minute * time.Duration(expMins)).Unix(),
		"salt": utils.RandNumberString(4),
	})
	return &Token{Token: token}
}

// Set sets the given key-val in to  Token.Claims
func (t *Token) Set(key, val string) *Token {
	t.Claims.(jwt.MapClaims)[key] = val
	return t
}

// SetMap add the given map into Token.Claims
func (t *Token) SetMap(m map[string]string) *Token {
	for k, v := range m {
		t.Set(k, v)
	}
	return t
}

// Sign retuns the SignedString with BaseSecretKey
func (t *Token) Sign() (string, error) {
	return t.SignedString([]byte(config.Config.BaseSecret))
}

// NewToString returns a jwt token string
func NewToString(name string, expMins int) (string, error) {
	return New(name, expMins).Sign()
}

// CheckSimple validates token
func CheckSimple(tkn string, name string) bool {
	return CheckWithVal(tkn, "name", name)
}

// CheckWithVal checks all Claims valus with the given key-val
func CheckWithVal(tkn, key, val string) bool {
	token, err := Parse(tkn)
	if err != nil || !token.Valid {
		return false
	}
	return token.Claims.(jwt.MapClaims)[key] == val
}

// CheckWithVals checks all Claims valus with the given map vals
func CheckWithVals(tkn string, vals map[string]string) bool {
	token, err := Parse(tkn)
	if err != nil || !token.Valid {
		return false
	}
	claims := token.Claims.(jwt.MapClaims)

	for k, v := range vals {
		if claims[k] != v {
			return false
		}
	}
	return true
}

// Parse parses the given tkn to a new Token
func Parse(tkn string) (*jwt.Token, error) {
	return jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config.BaseSecret), nil
	})
}

// ParseParams parses the tkn to the jwt.MapClaims
func ParseParams(tkn string) (jwt.MapClaims, error) {
	t, err := Parse(tkn)
	if err != nil {
		return nil, err
	}

	return t.Claims.(jwt.MapClaims), nil
}

// GetParam get param from the given  tkn
func GetParam(tkn, k string) string {
	m, err := ParseParams(tkn)
	if err != nil {
		return ""
	}

	return m[k].(string)
}
