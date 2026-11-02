package client

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
	createFn func(ctx context.Context, req CreateClientRequest) (*Client, error)
}

func (f *fakeService) Create(ctx context.Context, req CreateClientRequest) (*Client, error) {
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
		svc := &fakeService{
			createFn: func(ctx context.Context, req CreateClientRequest) (*Client, error) {
				return &Client{
					TenantID: req.TenantID,
					Name:     req.Name,
					Phone:    req.Phone,
				}, nil
			},
		}
		router := newTestRouter(NewHandler(svc))

		tenantID := uuid.New()
		body, _ := json.Marshal(CreateClientRequest{
			TenantID: tenantID,
			Name:     "John Doe",
			Phone:    "+5511999999999",
		})
		req := httptest.NewRequest(http.MethodPost, "/api/clients", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusCreated {
			t.Fatalf("expected status %d, got %d (body: %s)", http.StatusCreated, w.Code, w.Body.String())
		}

		var resp ClientResponse
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("failed to unmarshal response: %v", err)
		}
		if resp.Name != "John Doe" || resp.Phone != "+5511999999999" || resp.TenantID != tenantID.String() {
			t.Fatalf("unexpected response body: %+v", resp)
		}
	})

	t.Run("returns 400 when payload is invalid", func(t *testing.T) {
		svc := &fakeService{
			createFn: func(ctx context.Context, req CreateClientRequest) (*Client, error) {
				t.Fatalf("service should not be called with invalid payload")
				return nil, nil
			},
		}
		router := newTestRouter(NewHandler(svc))

		body, _ := json.Marshal(map[string]string{"name": "John Doe"})
		req := httptest.NewRequest(http.MethodPost, "/api/clients", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("expected status %d, got %d (body: %s)", http.StatusBadRequest, w.Code, w.Body.String())
		}
	})

	t.Run("returns 500 when service fails", func(t *testing.T) {
		svc := &fakeService{
			createFn: func(ctx context.Context, req CreateClientRequest) (*Client, error) {
				return nil, assertAnError
			},
		}
		router := newTestRouter(NewHandler(svc))

		body, _ := json.Marshal(CreateClientRequest{
			TenantID: uuid.New(),
			Name:     "John Doe",
			Phone:    "+5511999999999",
		})
		req := httptest.NewRequest(http.MethodPost, "/api/clients", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Fatalf("expected status %d, got %d (body: %s)", http.StatusInternalServerError, w.Code, w.Body.String())
		}
	})
}

var assertAnError = httpTestError("service failure")

type httpTestError string

func (e httpTestError) Error() string { return string(e) }
