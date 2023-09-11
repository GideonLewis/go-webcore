package user

import (
	"github.com/labstack/echo/v4"
	"github.com/megaqstar/web-core/payload"
	"github.com/megaqstar/web-core/presenter"
)

type IUseCase interface {
	Create(ctx echo.Context, req *payload.UserCreateRequest) error
	GetByID(ctx echo.Context, req *payload.UserGetByIDRequest) (*presenter.UserResponseWrapper, error)
	GetList(ctx echo.Context, req *payload.UserGetListRequest) (*presenter.ListUserResponseWrapper, error)
	Delete(ctx echo.Context, req *payload.UserDeleteRequest) error
	Update(ctx echo.Context, req *payload.UserUpdateRequest) (*presenter.UserResponseWrapper, error)
}
