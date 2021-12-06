package middleware

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.Use(middleware.Recover(),
		middleware.RequestID(),
		middleware.CORS(),
		middleware.Logger(),
	)
}
