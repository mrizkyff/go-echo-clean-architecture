package response

import (
	"github.com/google/uuid"
)

type UserResponseDto struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
	Phone    string    `json:"phone"`
	Address  string    `json:"address"`
	Age      int       `json:"age"`
}
