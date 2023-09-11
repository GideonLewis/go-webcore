package user

import (
	"github.com/labstack/echo/v4"
	"github.com/megaqstar/web-core/meg-pkg/common"
	"github.com/megaqstar/web-core/model"
)

type Repository interface {
	Create(ctx echo.Context, data *model.User) error
	Update(ctx echo.Context, data *model.User) error
	Delete(ctx echo.Context, data *model.User, unscoped bool) error
	GetByID(ctx echo.Context, id int64) (*model.User, error)
	GetList(
		ctx echo.Context,
		paginator common.Paginator,
		conditions interface{},
		orderBy string,
	) ([]model.User, error)
}
