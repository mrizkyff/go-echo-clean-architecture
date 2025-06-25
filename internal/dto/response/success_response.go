package response

import "github.com/labstack/echo/v4"

type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c echo.Context, status int, message string, data interface{}) error {
	httpStatus := GetStatusByCode(status)
	response := SuccessResponse{
		Status: httpStatus.Code,
		Message: func() string {
			if message != "" {
				return message
			}
			return httpStatus.Message
		}(),
		Data: data,
	}
	return c.JSON(200, response)
}
