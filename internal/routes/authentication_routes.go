package routes

import (
	"go-echo-clean-architecture/internal/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterAuthenticationRoute(handler *handlers.AuthenticationHandler, echo *echo.Echo) {
	echo.POST("/login", handler.Login)
	echo.POST("/register", handler.Register)
}
