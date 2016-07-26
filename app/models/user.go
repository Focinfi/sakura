package models

// User for users
type User struct {
	Model
	Name string `json:"name"`
}

// DisplayName implements Auth interface
func (user *User) DisplayName() string {
	return user.Name
}
