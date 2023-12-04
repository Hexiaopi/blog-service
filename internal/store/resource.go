package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
)

type ResourceStore interface {
	Create(ctx context.Context, article *entity.Resource) error
	Update(ctx context.Context, article *entity.Resource) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*entity.Resource, error)
	List(ctx context.Context, opt *entity.ListOption) ([]entity.Resource, error)
	Count(ctx context.Context, opt *entity.ListOption) (int64, error)
}
