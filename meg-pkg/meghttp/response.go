package meghttp

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type response struct{}

var Response response

func (r *response) Success(ctx echo.Context, data interface{}) error {
	return ctx.JSON(
		http.StatusOK,
		map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success",
			"data":    data,
		},
	)
}

func (r *response) Error(ctx echo.Context, err MegError) error {
	var errMessage string

	// if err.IsSentry {
	// 	teqlogger.GetLogger().Error(err.Raw.Error())
	// 	teqsentry.WithContext(&c).Error(err.Raw)
	// }

	if err.Raw != nil {
		errMessage = err.Raw.Error()
	}

	return ctx.JSON(err.HTTPCode, map[string]interface{}{
		"code":    err.ErrorCode,
		"message": err.Message,
		"info":    errMessage,
	})
}
