package user

import "context"

type Service interface {
	Create(ctx context.Context, req CreateUserRequest) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, req CreateUserRequest) (*User, error) {
	u := &User{
		TenantID: req.TenantID,
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
		Role: req.Role,
	}

	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}

	return u, nil
}
