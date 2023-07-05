package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/model"
)

type RoleStore interface {
	Get(ctx context.Context, id int) (*model.Role, error)
	Create(ctx context.Context, user *model.Role) error
	Update(ctx context.Context, user *model.Role) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, opt *model.ListOption) ([]model.Role, error)
	Count(ctx context.Context, opt *model.ListOption) (int64, error)
}
