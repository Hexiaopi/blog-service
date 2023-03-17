package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/model"
)

type TagStore interface {
	Create(ctx context.Context, tag *model.Tag) error
	Update(ctx context.Context, tag *model.Tag) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*model.Tag, error)
	List(ctx context.Context, opt *model.ListOption) ([]model.Tag, error)
	Count(ctx context.Context, opt *model.ListOption) (int64, error)
}
