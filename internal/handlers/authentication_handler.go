package handlers

import (
	"go-echo-clean-architecture/internal/dto/request"
	"go-echo-clean-architecture/internal/dto/response"
	"go-echo-clean-architecture/internal/models"
	"go-echo-clean-architecture/internal/services"

	"github.com/labstack/echo/v4"
)

type AuthenticationHandler struct {
	authenticationService services.AuthenticationService
	userService           services.UserService
}

func NewAuthenticationHandler(authenticationService services.AuthenticationService, userService services.UserService) *AuthenticationHandler {
	return &AuthenticationHandler{authenticationService: authenticationService, userService: userService}
}

func (receiver *AuthenticationHandler) Login(ctx echo.Context) error {
	// Get username & password from req body
	var loginRequest request.LoginRequest

	err := ctx.Bind(&loginRequest)
	if err != nil {
		return response.Error(ctx, 400, "Bad request")
	}

	accessToken, err := receiver.authenticationService.Login(loginRequest.Username, loginRequest.Password, ctx.Request().Context())
	if err != nil {
		return err
	}

	return response.Success(ctx, 200, "Success", accessToken)
}

func (receiver *AuthenticationHandler) Register(ctx echo.Context) error {
	user := &models.User{}
	if err := ctx.Bind(user); err != nil {
		return response.Error(ctx, 400, "invalid request")
	}
	createdUser, err := receiver.userService.CreateUser(user)
	if err != nil {
		return err
	}
	return response.Success(ctx, 200, "user created", createdUser)
}
