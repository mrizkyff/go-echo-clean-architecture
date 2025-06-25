package mappers

import (
	"go-echo-clean-architecture/internal/dto/response"
	"go-echo-clean-architecture/internal/models"
)

func MapToUserDto(user models.User) response.UserResponseDto {
	return response.UserResponseDto{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
		Phone:    user.Phone,
		Address:  user.Address,
		Age:      user.Age,
	}
}

func MapToUserDtos(users []*models.User) []response.UserResponseDto {
	usersResponseDtos := make([]response.UserResponseDto, len(users))
	for i, user := range users {
		usersResponseDtos[i] = MapToUserDto(*user)
	}
	return usersResponseDtos
}
