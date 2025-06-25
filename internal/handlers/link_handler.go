package handlers

import (
	"go-echo-clean-architecture/internal/dto/response"
	"go-echo-clean-architecture/internal/errors"
	"go-echo-clean-architecture/internal/models"
	"go-echo-clean-architecture/internal/services"
	"go-echo-clean-architecture/internal/utils"
	"go-echo-clean-architecture/pkg/rabbitmq/publisher"
	"math"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type LinkHandler struct {
	linkService          services.LinkService
	logActivityPublisher *publisher.ActivityLogPublisher
	jwtUtils             *utils.JWTUtil
}

func NewLinkHandler(linkService services.LinkService, logActivityPublisher *publisher.ActivityLogPublisher, jwtUtils *utils.JWTUtil) *LinkHandler {
	return &LinkHandler{linkService: linkService, logActivityPublisher: logActivityPublisher, jwtUtils: jwtUtils}
}

func (receiver *LinkHandler) DeleteLink(ctx echo.Context) error {
	// Get id from path params
	id := ctx.Param("id")

	linkId, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	link, err := receiver.linkService.DeleteLink(linkId)
	if err != nil {
		return err
	}

	return response.Success(ctx, 200, "success", link)
}

func (receiver *LinkHandler) UpdateLink(ctx echo.Context) error {
	// Get Id from path params
	id := ctx.Param("id")

	// Parse id to uuid
	linkId, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	// Get user info from ctx
	userFromContext, _ := receiver.jwtUtils.GetUserFromContext(ctx)

	// Get req body
	linkRequest := &models.Link{}
	linkRequest.UpdatedBy = userFromContext.ID

	// Parse req body to model
	err = ctx.Bind(linkRequest)
	if err != nil {
		return response.Error(ctx, 404, "invalid request")
	}

	// Update link
	updatedLink, err := receiver.linkService.UpdateLink(linkId, linkRequest)
	if err != nil {
		return err
	}

	return response.Success(ctx, 200, "success", updatedLink)
}

func (receiver *LinkHandler) CreateLink(ctx echo.Context) error {
	// Create variable for store link request body
	linkRequestBody := &models.Link{}

	// Get user info from ctx
	userFromContext, _ := receiver.jwtUtils.GetUserFromContext(ctx)
	linkRequestBody.CreatedBy = userFromContext.ID

	// Parse to linkRequestBody
	err := ctx.Bind(linkRequestBody)
	if err != nil {
		return response.Error(ctx, 400, "invalid request")
	}

	createdUser, err := receiver.linkService.CreateLink(linkRequestBody)
	if err != nil {
		return err
	}

	return response.Success(ctx, 200, "success", createdUser)
}

func (receiver *LinkHandler) GetAllLinks(ctx echo.Context) error {
	pagination := *utils.NewPaginationType()
	pagination.BindQueryParams(ctx)

	links, totalRecords, err := receiver.linkService.GetAllLinkWithPagination(pagination)
	if err != nil {
		return err
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

	return response.SuccessWithPagination(ctx, links, meta)
}

func (receiver *LinkHandler) GetLinkById(ctx echo.Context) error {
	idStr := ctx.Param("id")

	// Parse to uuid
	id, err := uuid.Parse(idStr)
	if err != nil {
		return errors.NewBadRequestError("Invalid link ID format")
	}

	// Get user info from ctx
	userFromContext, _ := receiver.jwtUtils.GetUserFromContext(ctx)

	// Call handler
	link, err := receiver.linkService.GetLinkById(id, userFromContext.ID)
	if err != nil {
		return errors.NewNotFoundError("Link not found")
	}

	// Publish activity log
	receiver.logActivityPublisher.Publish(ctx, id, userFromContext.ID)

	return response.Success(ctx, 200, "success", link)
}
