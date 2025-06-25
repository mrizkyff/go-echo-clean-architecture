package middlewares

import (
	"fmt"
	"net/http"
	"runtime"

	"go-echo-clean-architecture/internal/dto/response"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// RecoverMiddleware returns a middleware that recovers from panics
func RecoverMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}

					// Collect stack trace
					stack := make([]byte, 4<<10) // 4KB
					length := runtime.Stack(stack, false)
					stackTrace := string(stack[:length])

					// Log the error and stack trace
					log.Errorf("[PANIC RECOVERED] %v %s\n%s", err, c.Request().URL, stackTrace)

					// Return a 500 error
					if !c.Response().Committed {
						_ = response.Error(c, http.StatusInternalServerError, "Server error occurred")
					}
				}
			}()

			return next(c)
		}
	}
}
