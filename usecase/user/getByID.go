package user

import (
	"github.com/labstack/echo/v4"
	"github.com/megaqstar/web-core/payload"
	"github.com/megaqstar/web-core/presenter"
)

func (u *UseCase) GetByID(
	ctx echo.Context,
	req *payload.UserGetByIDRequest) (
	*presenter.UserResponseWrapper,
	error) {

	user, err := u.UserRepo.GetByID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	return &presenter.UserResponseWrapper{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Gender:    user.Gender,
		Age:       user.Age,
		Phone:     user.Phone,
		Email:     user.Email,
		Country:   user.Country,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
