package user

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/megaqstar/web-core/meg-pkg/meghttp"
	"github.com/megaqstar/web-core/payload"
	"gorm.io/gorm"
)

func (r *Route) delete(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return meghttp.Response.Error(ctx, meghttp.ErrParseInt(err, "user"))
	}

	err = r.useCase.User.Delete(ctx, &payload.UserDeleteRequest{UserID: id})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return meghttp.Response.Error(ctx, meghttp.ErrNotFound(err, "user"))
		}
		return meghttp.Response.Error(ctx, meghttp.ErrDelete(err, "user"))
	}

	return meghttp.Response.Success(ctx, nil)
}
