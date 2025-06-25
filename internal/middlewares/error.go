package middlewares

import (
	stderrors "errors"
	"fmt"
	"go-echo-clean-architecture/internal/errors"
	"net/http"

	"go-echo-clean-architecture/internal/dto/response"

	"github.com/labstack/echo/v4"
)

// ErrorHandlingMiddleware provides a centralized error handling mechanism
func ErrorHandlingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err == nil {
				return nil
			}

			if c.Response().Committed {
				return err
			}

			// Handle app errors
			var appErr *errors.AppError
			if stderrors.As(err, &appErr) {
				return response.Error(c, appErr.Code, appErr.Message)
			}

			// Handle other error types
			switch e := err.(type) {
			case *echo.HTTPError:
				return response.Error(c, e.Code, fmt.Sprintf("%v", e.Message))
			case *echo.BindingError:
				return response.Error(c, http.StatusBadRequest, "Invalid request data: "+e.Error())
			default:
				return response.Error(c, http.StatusInternalServerError, "An internal server error occurred")
			}
		}
	}
}
