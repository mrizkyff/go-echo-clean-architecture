package middlewares

import (
	"go-echo-clean-architecture/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	jwt *utils.JWTUtil
}

func NewAuthMiddleware(jwt *utils.JWTUtil) *AuthMiddleware {
	return &AuthMiddleware{jwt: jwt}
}

func (a *AuthMiddleware) Authenticate() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Ekstrak token dari header
			tokenString, err := utils.ExtractTokenFromHeader(c.Request())
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized: " + err.Error()})
			}

			// Validasi token
			claims, err := a.jwt.ValidateToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized: " + err.Error()})
			}

			// Set claims ke context untuk digunakan di handler berikutnya
			c.Set("userID", claims.ID)
			c.Set("username", claims.Username)
			c.Set("email", claims.Email)
			c.Set("claims", claims)

			return next(c)
		}
	}
}
