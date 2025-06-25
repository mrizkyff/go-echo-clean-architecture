package utils

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

type PaginationType struct {
	Page     int
	PageSize int
	SortBy   string
	SortDir  string
}

func NewPaginationType() *PaginationType {
	return &PaginationType{Page: 0, PageSize: 10, SortBy: "created_at", SortDir: "desc"}
}

func (receiver *PaginationType) BindQueryParams(ctx echo.Context) {
	if pageParam := ctx.QueryParam("page"); pageParam != "" {
		if pageInt, err := strconv.Atoi(pageParam); err == nil && pageInt > 0 {
			receiver.Page = pageInt
		}
	}

	if pageSizeParam := ctx.QueryParam("pageSize"); pageSizeParam != "" {
		if pageSizeInt, err := strconv.Atoi(pageSizeParam); err == nil && pageSizeInt > 0 {
			receiver.PageSize = pageSizeInt
		}
	}

	if sortByParam := ctx.QueryParam("sortBy"); sortByParam != "" {
		receiver.SortBy = sortByParam
	}

	if sortDirParam := ctx.QueryParam("sortDir"); sortDirParam != "" {
		if sortDirParam == "asc" || sortDirParam == "desc" {
			receiver.SortDir = sortDirParam
		}
	}
}
