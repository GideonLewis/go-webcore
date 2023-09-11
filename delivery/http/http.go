package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/megaqstar/web-core/delivery/http/user"
	"github.com/megaqstar/web-core/usecase"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func HTTPServe(usecase *usecase.UseCase) *echo.Echo {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
	}))

	// API documentation: http://localhost:8000/docs/index.html
	e.GET("/docs/*", echoSwagger.WrapHandler)

	userGroup := e.Group("v1/user")
	user.Init(userGroup, usecase)

	return e
}
