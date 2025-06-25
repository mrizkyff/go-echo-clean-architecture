package services

import (
	"context"
	"go-echo-clean-architecture/internal/errors"
	"go-echo-clean-architecture/internal/models"
	"go-echo-clean-architecture/internal/repositories"
	"go-echo-clean-architecture/internal/utils"

	"github.com/google/uuid"
)

type AuthenticationService interface {
	Login(username string, password string, ctx context.Context) (string, error)
	Register(user *models.User, ctx context.Context) (*models.User, error)
	Logout(id uuid.UUID) error
}

type AuthenticationServiceImpl struct {
	userRepository      repositories.UserRepository
	userService         UserService
	jwtUtils            *utils.JWTUtil
	userRedisRepository repositories.UserRedisRepository
}

func NewAuthenticationServiceImpl(userRepository repositories.UserRepository, userService UserService, jwtUtils *utils.JWTUtil, userRedisRepository repositories.UserRedisRepository) *AuthenticationServiceImpl {
	return &AuthenticationServiceImpl{userRepository: userRepository, userService: userService, jwtUtils: jwtUtils, userRedisRepository: userRedisRepository}
}

func (a *AuthenticationServiceImpl) Login(username string, password string, ctx context.Context) (string, error) {
	// Create context for redis
	backgroundCtx := context.Background()

	// Find user by username or email or phone on redis cache
	user := a.userRedisRepository.FindByUsernameOrEmailOrPhone(username, backgroundCtx)

	// If user not found in redis cache, find on database
	if user == nil {
		var err error
		// Find user by username or email or phone
		user, err = a.userRepository.FindByUserNameOrEmailOrPhone(username, ctx)
		if (user == nil) || (err != nil) {
			return "", errors.NewUnauthorizedError("Username or password not match1")
		}
	}

	// Comparing password
	verifiedPassword, err := utils.VerifyPassword(password, user.Password)
	if err != nil {
		return "", errors.NewUnauthorizedError("Username or password not match4")
	}
	if verifiedPassword != true {
		return "", errors.NewUnauthorizedError("Username or password not match5")
	}

	// Generating jwt
	accessToken, err := a.jwtUtils.GenerateToken(user.ID.String(), user.Username, user.Email)
	if err != nil {
		return "", errors.NewInternalServerError("Failed to generate jwt token")
	}

	// Store user in redis as cache
	go a.userRedisRepository.Update(user, user, backgroundCtx)
	return accessToken, nil
}

func (a *AuthenticationServiceImpl) Register(user *models.User, ctx context.Context) (*models.User, error) {
	createUser, err := a.userService.CreateUser(user)
	if err != nil {
		return nil, err
	}

	// Store user in redis as cache
	a.userRedisRepository.Create(ctx, user)

	return createUser, nil
}

func (a *AuthenticationServiceImpl) Logout(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
