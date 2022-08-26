package repos

import (
	"git.wildberries.ru/branch-office/go-services/portal-api/logger"
	"git.wildberries.ru/branch-office/go-services/portal-api/provider"
)

type Repositories struct {
	UserDBRepo UserDBRepo
}

func New(p provider.Provider, l logger.Logger) *Repositories {
	UserDBRepo := N NewUserDBRepo(p, l)
	return &Repositories{
		UserDBRepo
	}
}
