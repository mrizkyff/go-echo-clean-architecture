package routes

import (
	"go-echo-clean-architecture/internal/handlers"
	"go-echo-clean-architecture/internal/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(handler *handlers.UserHandler, echo *echo.Echo, authentication *middlewares.AuthMiddleware) {
	echo.GET("/users", handler.GetAllUsers, authentication.Authenticate())
	echo.GET("/users/:id", handler.GetUserByID, authentication.Authenticate())
	echo.POST("/users", handler.CreateUser, authentication.Authenticate())
	echo.PUT("/users/:id", handler.UpdateUser, authentication.Authenticate())
	echo.DELETE("/users/:id", handler.DeleteUser, authentication.Authenticate())
}
