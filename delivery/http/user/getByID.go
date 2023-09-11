package user

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/megaqstar/web-core/meg-pkg/meghttp"
	"github.com/megaqstar/web-core/payload"
	"gorm.io/gorm"
)

// GetByID user by id
// @Summary Get an user
// @Description Get user by id
// @Tags user
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path int true "id"
// @Success 200 {object} presenter.UserResponseWrapper
// @Router /user/{id} [get] .
func (r *Route) getByID(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return meghttp.Response.Error(ctx, meghttp.ErrParseInt(err, "user"))
	}
	resp, err := r.useCase.User.GetByID(ctx, &payload.UserGetByIDRequest{UserID: id})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return meghttp.Response.Error(ctx, meghttp.ErrNotFound(err, "user"))
		}
		return meghttp.Response.Error(ctx, meghttp.ErrGet(err, "user"))
	}

	return meghttp.Response.Success(ctx, resp)
}
