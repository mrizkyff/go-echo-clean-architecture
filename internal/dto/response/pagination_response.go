package response

import "github.com/labstack/echo/v4"

type PaginationMeta struct {
	CurrentPage  int   `json:"current_page"`
	PerPage      int   `json:"per_page"`
	TotalPages   int   `json:"total_pages"`
	TotalRecords int64 `json:"total_records"`
}

type PaginationResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    interface{}    `json:"data,omitempty"`
	Meta    PaginationMeta `json:"meta"`
}

func SuccessWithPagination(c echo.Context, data interface{}, meta PaginationMeta) error {
	httpStatus := GetStatusByCode(200)
	response := PaginationResponse{
		Status:  httpStatus.Code,
		Message: httpStatus.Message,
		Data:    data,
		Meta:    meta,
	}
	return c.JSON(200, response)
}
