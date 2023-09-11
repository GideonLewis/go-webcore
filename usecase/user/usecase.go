package user

import (
	"github.com/megaqstar/web-core/repository"
	"github.com/megaqstar/web-core/repository/user"
)

type UseCase struct {
	UserRepo user.Repository
}

func NewUseCase(repo *repository.Repository) IUseCase {
	return &UseCase{
		UserRepo: repo.User,
	}
}
