package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/model"
)

type OperationStore interface {
	Create(ctx context.Context, log *entity.OperationLog) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*entity.OperationLog, error)
	List(ctx context.Context, opt *entity.ListOption) ([]model.OperationLog, error)
	Count(ctx context.Context, opt *entity.ListOption) (int64, error)
}
