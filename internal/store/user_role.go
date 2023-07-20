package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/model"
)

type UserRoleStore interface {
	Create(ctx context.Context, userRole *model.UserRole) error
	Delete(ctx context.Context, userRole *model.UserRole) error
}