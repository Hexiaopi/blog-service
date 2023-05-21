package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/model"
)

type ResourceStore interface {
	Create(ctx context.Context, article *model.Resource) error
	Update(ctx context.Context, article *model.Resource) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*model.Resource, error)
	List(ctx context.Context, opt *model.ListOption) ([]model.Resource, error)
	Count(ctx context.Context, opt *model.ListOption) (int64, error)
}
