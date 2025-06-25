package routes

import (
	"go-echo-clean-architecture/internal/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterAccessLogRoutes(handler *handlers.AccessLogHandler, echo *echo.Echo) {
	echo.GET("/access-log", handler.GetAllLogHandler)
}
