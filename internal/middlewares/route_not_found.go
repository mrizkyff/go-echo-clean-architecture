package middlewares

import (
	"errors"
	"net/http"

	"go-echo-clean-architecture/internal/dto/response"

	"github.com/labstack/echo/v4"
)

// NotFoundHandler returns a middleware function that handles routes that aren't found
func NotFoundHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Execute the next middleware/handler in the chain
			err := next(c)

			// If route is not found (404)
			if err != nil && errors.Is(err, echo.ErrNotFound) {
				return response.Error(c, http.StatusNotFound, "Route not found")
			}

			return err
		}
	}
}

// RegisterNotFoundRoute registers a catch-all route to handle 404s
func RegisterNotFoundRoute(e *echo.Echo) {
	e.Any("/*", func(c echo.Context) error {
		return response.Error(c, http.StatusNotFound, "Route not found")
	})
}
