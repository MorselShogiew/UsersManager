package usecases

import (
	"git.wildberries.ru/branch-office/go-services/portal-api/logger"
	"git.wildberries.ru/branch-office/go-services/portal-api/repos"
)

type UserService struct {
	userDBRepo repos.UserDBRepo

	l logger.Logger
}

func New(r *repos.Repositories, l logger.Logger) *UserService {
	return &UserService{
		r.UserDBRepo,
		l,
	}
}
