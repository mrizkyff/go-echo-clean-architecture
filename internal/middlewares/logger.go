package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// LoggerMiddleware returns a middleware that logs HTTP requests using logrus
func LoggerMiddleware() echo.MiddlewareFunc {
	// Initialize logrus
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Record start time
			start := time.Now()

			// Process request
			err := next(c)

			// Calculate response time
			stop := time.Now()
			latency := stop.Sub(start)

			// Get request and response details
			req := c.Request()
			res := c.Response()
			status := res.Status

			// Create log fields
			fields := logrus.Fields{
				"method":     req.Method,
				"path":       req.URL.Path,
				"status":     status,
				"latency":    latency,
				"ip":         c.RealIP(),
				"user_agent": req.UserAgent(),
			}

			// Log with appropriate level based on status code
			if status >= 500 {
				log.WithFields(fields).Error("Server error")
			} else if status >= 400 {
				log.WithFields(fields).Warn("Client error")
			} else {
				log.WithFields(fields).Info("Request completed")
			}

			return err
		}
	}
}
