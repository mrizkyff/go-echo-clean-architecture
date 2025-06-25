package routes

import (
	"go-echo-clean-architecture/internal/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterHelloRoutes(handler *handlers.HelloHandler, echo *echo.Echo) {
	echo.GET("/hello", handler.Hello)
}
