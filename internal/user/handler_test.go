package user

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type fakeService struct {
	createFn func(ctx context.Context, req CreateUserRequest) (*User, error)
}

func (f *fakeService) Create(ctx context.Context, req CreateUserRequest) (*User, error) {
	return f.createFn(ctx, req)
}

func newTestRouter(h *Handler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h.RegisterRoutes(r.Group("/api"))
	return r
}

func TestHandler_Create(t *testing.T) {
	t.Run("returns 201 with the created client", func(t *testing.T) {

	})
}
