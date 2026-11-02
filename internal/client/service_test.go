package client

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

var testTenantID = uuid.New()

type fakeRepository struct {
	createFn func(ctx context.Context, c *Client) error
}

func (f *fakeRepository) Create(ctx context.Context, c *Client) error {
	return f.createFn(ctx, c)
}

func TestService_Create(t *testing.T) {
	t.Run("creates a client with the request data", func(t *testing.T) {
		var captured *Client
		repo := &fakeRepository{
			createFn: func(ctx context.Context, c *Client) error {
				captured = c
				return nil
			},
		}
		svc := NewService(repo)

		got, err := svc.Create(context.Background(), CreateClientRequest{
			TenantID: testTenantID,
			Name:     "John Doe",
			Phone:    "+5511999999999",
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if got != captured {
			t.Fatalf("expected service to return the created client")
		}
		if got.TenantID != testTenantID || got.Name != "John Doe" || got.Phone != "+5511999999999" {
			t.Fatalf("unexpected client data: %+v", got)
		}
	})

	t.Run("propagates repository errors", func(t *testing.T) {
		repoErr := errors.New("db error")
		repo := &fakeRepository{
			createFn: func(ctx context.Context, c *Client) error {
				return repoErr
			},
		}
		svc := NewService(repo)

		got, err := svc.Create(context.Background(), CreateClientRequest{
			TenantID: testTenantID,
			Name:     "John Doe",
			Phone:    "+5511999999999",
		})

		if !errors.Is(err, repoErr) {
			t.Fatalf("expected repository error, got %v", err)
		}
		if got != nil {
			t.Fatalf("expected nil client on error, got %+v", got)
		}
	})
}
