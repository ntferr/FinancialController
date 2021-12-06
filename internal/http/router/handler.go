package http

import (
	"FinancialController/internal/controller/health"
	"FinancialController/internal/http/middleware"

	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()

	middleware.Init(e)

	h := health.NewHealth()

	router := e.Group("financial/v1")
	{
		router.GET("/health", h.Status)
	}
}
