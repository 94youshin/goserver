package v1

import "github.com/go-playground/validator/v10"

type User struct {
	BaseModel

	// Required: true
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`

	Password string `json:"password"`

	NickName string `json:"nickname"`

	Email string `json:"email"`
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
