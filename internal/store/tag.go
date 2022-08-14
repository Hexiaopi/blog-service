package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/model"
)

type TagStore interface {
	Create(ctx context.Context, tag *entity.Tag) error
	Update(ctx context.Context, tag *entity.Tag) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*model.Tag, error)
	List(ctx context.Context, opt *entity.ListOption) ([]model.Tag, int64, error)
}
