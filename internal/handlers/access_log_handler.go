package handlers

import (
	"go-echo-clean-architecture/internal/dto/response"
	"go-echo-clean-architecture/internal/mappers"
	"go-echo-clean-architecture/internal/services"
	"go-echo-clean-architecture/internal/utils"
	"math"

	"github.com/labstack/echo/v4"
)

type AccessLogHandler struct {
	accessLogService services.AccessLogService
}

func NewAccessLogHandler(accessLogService services.AccessLogService) *AccessLogHandler {
	return &AccessLogHandler{accessLogService: accessLogService}
}

func (receiver *AccessLogHandler) GetAllLogHandler(ctx echo.Context) error {
	// Get pagination parameters from query string
	pagination := *utils.NewPaginationType()
	// Parse page parameter if provided
	pagination.BindQueryParams(ctx)

	accessLogs, totalRecords, err := receiver.accessLogService.GetAll(pagination)
	if err != nil {
		return err
	}

	accessLogsDtos := make([]*response.AccessLogResponseDto, len(accessLogs))
	for i, accessLog := range accessLogs {
		accessLogsDtos[i] = mappers.MapAccessLogToDto(accessLog)
	}

	// Count total pages
	totalPages := int(math.Ceil(float64(totalRecords) / float64(pagination.PageSize)))

	// Create pagination metadata
	meta := response.PaginationMeta{
		CurrentPage:  pagination.Page,
		PerPage:      pagination.PageSize,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
	}

	return response.SuccessWithPagination(ctx, accessLogsDtos, meta)
}
