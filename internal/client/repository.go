package client

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, client *Client) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, client *Client) error {
	return r.db.WithContext(ctx).Create(client).Error
}
