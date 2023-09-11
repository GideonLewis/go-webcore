package user

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/megaqstar/web-core/meg-pkg/meghttp"
	"github.com/megaqstar/web-core/payload"
)

func (r *Route) update(ctx echo.Context) error {
	idStr := ctx.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return meghttp.Response.Error(ctx, meghttp.ErrInvalidParams(err, "user"))
	}

	req := payload.UserUpdateRequest{
		ID: id,
	}

	if err = ctx.Bind(&req); err != nil {
		return meghttp.Response.Error(ctx, meghttp.ErrInvalidParams(err, "user"))
	}

	resp, err := r.useCase.User.Update(ctx, &req)
	if err != nil {
		return meghttp.Response.Error(ctx, meghttp.ErrUpdate(err, "user"))
	}

	return meghttp.Response.Success(ctx, resp)
}
