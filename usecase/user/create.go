package user

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/megaqstar/web-core/meg-pkg/common"
	"github.com/megaqstar/web-core/model"
	"github.com/megaqstar/web-core/payload"
)

func (u *UseCase) validateCreate(
	ctx echo.Context,
	req *payload.UserCreateRequest) error {

	if !common.IsGenderValid(req.Gender) {
		return errors.New("gender invalid")
	}

	if ok, _ := common.IsEmailValid(req.Email); !ok {
		return errors.New("email invalid")
	}

	if req.Age <= 0 {
		return errors.New("age invalid")
	}

	return nil
}

func (u *UseCase) Create(
	ctx echo.Context,
	req *payload.UserCreateRequest) error {

	if err := u.validateCreate(ctx, req); err != nil {
		return err
	}

	userReq := &model.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Gender:    req.Gender,
		Age:       req.Age,
		Phone:     req.Phone,
		Email:     req.Email,
		Country:   req.Country,
	}

	if err := u.UserRepo.Create(ctx, userReq); err != nil {
		return err
	}

	return nil
}
