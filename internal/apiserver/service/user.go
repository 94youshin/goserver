package service

import (
	"context"

	v1 "github.com/usmhk/goserver/internal/apiserver/model/v1"
	"github.com/usmhk/goserver/internal/apiserver/store"
	metav1 "github.com/usmhk/goserver/pkg/meta/v1"
)

type UserSrv interface {
	Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error
}

type userService struct {
	store store.Factory
}

func newUsers(srv *service) *userService {
	return userService{
		store: srv.store,
	}
}

func (u *userService) Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error {
	return nil
}
