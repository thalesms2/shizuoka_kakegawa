package client

import "context"

type Service interface {
	Create(ctx context.Context, req CreateClientRequest) (*Client, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, req CreateClientRequest) (*Client, error) {
	c := &Client{
		TenantID: req.TenantID,
		Name:     req.Name,
		Phone:    req.Phone,
	}

	if err := s.repo.Create(ctx, c); err != nil {
		return nil, err
	}

	return c, nil
}
