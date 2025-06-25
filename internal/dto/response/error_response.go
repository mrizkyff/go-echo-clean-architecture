package response

import "github.com/labstack/echo/v4"

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

func Error(c echo.Context, errorCode int, message string) error {
	httpStatus := GetStatusByCode(errorCode)
	response := ErrorResponse{
		Status: httpStatus.Code,
		Message: func() string {
			if message != "" {
				return message
			}
			return httpStatus.Message
		}(),
	}
	return c.JSON(errorCode, response)
}
