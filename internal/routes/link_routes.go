package routes

import (
	"go-echo-clean-architecture/internal/handlers"
	"go-echo-clean-architecture/internal/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterLinkRoutes(handler *handlers.LinkHandler, echo *echo.Echo, authentication *middlewares.AuthMiddleware) {
	echo.GET("/links", handler.GetAllLinks, authentication.Authenticate())
	echo.GET("/links/:id", handler.GetLinkById, authentication.Authenticate())
	echo.POST("/links", handler.CreateLink, authentication.Authenticate())
	echo.PUT("/links/:id", handler.UpdateLink, authentication.Authenticate())
	echo.DELETE("/links/:id", handler.DeleteLink, authentication.Authenticate())
}
