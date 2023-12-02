package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
)

type RoleStore interface {
	Get(ctx context.Context, id int) (*entity.Role, error)
	Create(ctx context.Context, user *entity.Role) error
	Update(ctx context.Context, user *entity.Role) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, opt *entity.ListOption) ([]entity.Role, error)
	Count(ctx context.Context, opt *entity.ListOption) (int64, error)
}
