package usecases

import (
	"github.com/MorselShogiew/UsersManager/repos"
)

type UserService struct {
	userDBRepo repos.UserDBRepo
}

func New(r *repos.Repositories) *UserService {
	return &UserService{
		r.UserDBRepo,
	}
}
