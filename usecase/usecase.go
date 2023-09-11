package usecase

import (
	"github.com/megaqstar/web-core/repository"
	"github.com/megaqstar/web-core/usecase/user"
)

type UseCase struct {
	User user.IUseCase
}

func NewUseCase(repo *repository.Repository) *UseCase {
	return &UseCase{
		User: user.NewUseCase(repo),
	}
}
