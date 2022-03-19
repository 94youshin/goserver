package v1

import "github.com/go-playground/validator/v10"

type User struct {
	BaseModel

	// Required: true
	Username string `json:"username"`
}

func (u *User) TableName() string {
	return "user"
}

type UserList struct {
	Items []*User `json:"items"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
