package user

import "github.com/usmhk/goserver/internal/apiserver/store"

type UserController struct {
}

func NewUserController(store store.Factory) *UserController {
	return &UserController{}
}
