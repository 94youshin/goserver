package service

import "github.com/usmhk/goserver/internal/apiserver/store"

type Service interface {
	Users() UserSrv
}

type service struct {
	store store.Factory
}

func NewService(store store.Factory) Service {
	return &service{store: sotre}
}

func (s *service) Users() UserSrv {
	return newUsers(s)
}
