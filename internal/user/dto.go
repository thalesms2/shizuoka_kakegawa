package user

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	TenantID uuid.UUID `json:"tenant_id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role string `json:"role" binding:"required"`
}

type UserResponse struct {
	ID string `json:"id"`
	TenantID string `json:"tenant_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(u *User) UserResponse {
	return UserResponse{
		ID: u.ID.String(),
		TenantID: u.TenantID.String(),
		Name: u.Name,
		Email: u.Email,
		Role: u.Role,
		CreatedAt: u.CreatedAt,
	}
}
