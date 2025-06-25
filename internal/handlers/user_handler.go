package handlers

import (
	"go-echo-clean-architecture/internal/dto/response"
	"go-echo-clean-architecture/internal/mappers"
	"go-echo-clean-architecture/internal/models"
	"go-echo-clean-architecture/internal/services"
	"go-echo-clean-architecture/internal/utils"
	"math"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService services.UserService
	jwtUtils    *utils.JWTUtil
}

func NewUserHandler(userService services.UserService, jwtUtils *utils.JWTUtil) *UserHandler {
	return &UserHandler{userService: userService, jwtUtils: jwtUtils}
}

func (h *UserHandler) Test(c echo.Context) error {
	x, err := h.userService.GetAllUsers()
	if err != nil {
		return response.Error(c, 500, "failed to get all users")
	}
	return response.Success(c, 200, "success", x)
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	// Get pagination parameters from query string
	pagination := *utils.NewPaginationType()
	// Parse page parameter if provided
	pagination.BindQueryParams(c)

	// Get users with pagination and sorting
	users, totalRecords, err := h.userService.GetAllUsersWithPagination(pagination)
	if err != nil {
		return err
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(totalRecords) / float64(pagination.PageSize)))

	// Create pagination metadata
	meta := response.PaginationMeta{
		CurrentPage:  pagination.Page,
		PerPage:      pagination.PageSize,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
	}

	return response.SuccessWithPagination(c, mappers.MapToUserDtos(users), meta)
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	// Get user info from ctx
	userFromContext, _ := h.jwtUtils.GetUserFromContext(c)

	user, err := h.userService.GetUserByID(userFromContext.ID)
	if err != nil {
		return err
	}
	return response.Success(c, 200, "success", mappers.MapToUserDto(*user))
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return response.Error(c, 400, "invalid request")
	}

	// Get user info from ctx
	userFromContext, _ := h.jwtUtils.GetUserFromContext(c)
	user.CreatedBy = userFromContext.ID

	createdUser, err := h.userService.CreateUser(user)
	if err != nil {
		return err
	}
	return response.Success(c, 200, "user created", mappers.MapToUserDto(*createdUser))
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	user := &models.User{}

	// Get user info from ctx
	userFromContext, _ := h.jwtUtils.GetUserFromContext(c)
	user.UpdatedBy = userFromContext.ID

	if err := c.Bind(user); err != nil {
		return response.Error(c, 400, "invalid request")
	}
	updatedUser, err := h.userService.UpdateUser(userFromContext.ID, user)
	if err != nil {
		return err
	}
	return response.Success(c, 200, "user updated", mappers.MapToUserDto(*updatedUser))
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	// Get user info from ctx
	userFromContext, _ := h.jwtUtils.GetUserFromContext(c)

	err := h.userService.DeleteUser(userFromContext.ID, userFromContext.ID)
	if err != nil {
		return err
	}
	return response.Success(c, 200, "user deleted", nil)
}
