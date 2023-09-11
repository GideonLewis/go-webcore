package user

import (
	"github.com/labstack/echo/v4"

	"github.com/megaqstar/web-core/meg-pkg/meghttp"
	"github.com/megaqstar/web-core/payload"
)

// Create User
// @Summary Create User
// @Description Create a User
// @Tags User
// @Accept  json
// @Produce json
// @Security AuthToken
// @Param req body payload.CreateUserRequest true "User info"
// @Success 200 {object} presenter.UserResponseWrapper
// @Router /users [post] .
func (r *Route) create(ctx echo.Context) error {
	var (
		req = payload.UserCreateRequest{}
	)
	if err := ctx.Bind(&req); err != nil {
		return meghttp.Response.Error(ctx, meghttp.ErrBind(err, "user"))
	}

	err := r.useCase.User.Create(ctx, &req)
	if err != nil {
		return meghttp.Response.Error(ctx, meghttp.ErrCreate(err, "user"))
	}

	return meghttp.Response.Success(ctx, nil)
}
