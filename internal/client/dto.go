package client

import (
	"time"

	"github.com/google/uuid"
)

type CreateClientRequest struct {
	TenantID uuid.UUID `json:"tenant_id" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Phone    string    `json:"phone" binding:"required"`
}

type ClientResponse struct {
	ID        string    `json:"id"`
	TenantID  string    `json:"tenant_id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}

func newClientResponse(c *Client) ClientResponse {
	return ClientResponse{
		ID:        c.ID.String(),
		TenantID:  c.TenantID.String(),
		Name:      c.Name,
		Phone:     c.Phone,
		CreatedAt: c.CreatedAt,
	}
}
