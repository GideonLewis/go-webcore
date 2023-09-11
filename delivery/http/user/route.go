package user

import (
	"github.com/labstack/echo/v4"
	"github.com/megaqstar/web-core/usecase"
)

type Route struct {
	useCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{useCase: useCase}

	group.POST("", r.create)
	group.GET("/:id", r.getByID)
	group.GET("", r.getList)
	group.DELETE("/:id", r.delete)
	group.PUT("/:id", r.update)
}
