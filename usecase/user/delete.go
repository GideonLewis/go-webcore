package user

import (
	"github.com/labstack/echo/v4"
	"github.com/megaqstar/web-core/payload"
)

func (u *UseCase) Delete(ctx echo.Context, req *payload.UserDeleteRequest) error {
	user, err := u.UserRepo.GetByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	err = u.UserRepo.Delete(ctx, user, false)
	if err != nil {
		return err
	}

	return nil
}
