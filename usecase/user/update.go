package user

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/megaqstar/web-core/meg-pkg/common"
	"github.com/megaqstar/web-core/model"
	"github.com/megaqstar/web-core/payload"
	"github.com/megaqstar/web-core/presenter"
)

func (u *UseCase) validateUpdate(ctx echo.Context, req *payload.UserUpdateRequest) (*model.User, error) {
	user, err := u.UserRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	if !common.IsGenderValid(req.Gender) {
		return nil, errors.New("gender invalid")
	}

	if ok, _ := common.IsEmailValid(req.Email); !ok {
		return nil, errors.New("email invalid")
	}

	if req.Age <= 0 {
		return nil, errors.New("age invalid")
	}

	return user, nil
}

func (u *UseCase) Update(ctx echo.Context, req *payload.UserUpdateRequest) (*presenter.UserResponseWrapper, error) {
	user, err := u.validateUpdate(ctx, req)
	if err != nil {
		return nil, err
	}

	user.FirstName = req.FirstName

	err = u.UserRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &presenter.UserResponseWrapper{}, nil
}
