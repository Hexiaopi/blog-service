package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
)

type UserRoleStore interface {
	Create(ctx context.Context, userId, roleId int) error
	Delete(ctx context.Context, userId, roleId int) error
	ListUserRole(ctx context.Context, userId int) ([]entity.Role, error)
}
