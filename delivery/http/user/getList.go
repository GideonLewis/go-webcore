package user

import (
	"github.com/labstack/echo/v4"
	"github.com/megaqstar/web-core/meg-pkg/meghttp"
	"github.com/megaqstar/web-core/payload"
)

func (r *Route) getList(ctx echo.Context) error {
	var (
		req = payload.UserGetListRequest{}
	)

	if err := ctx.Bind(&req); err != nil {
		return meghttp.Response.Error(ctx, meghttp.ErrInvalidParams(err, "user"))
	}

	resp, err := r.useCase.User.GetList(ctx, &req)
	if err != nil {
		return meghttp.Response.Error(ctx, meghttp.ErrGet(err, "user"))
	}

	return meghttp.Response.Success(ctx, resp)
}
