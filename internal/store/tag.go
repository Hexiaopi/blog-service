package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
)

type TagStore interface {
	Create(ctx context.Context, tag *entity.Tag) error
	Update(ctx context.Context, tag *entity.Tag) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*entity.Tag, error)
	List(ctx context.Context, opt *entity.ListOption) ([]entity.Tag, error)
	Count(ctx context.Context, opt *entity.ListOption) (int64, error)
}
