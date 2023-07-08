package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/model"
)

type OperationStore interface {
	Create(ctx context.Context, log *model.OperationLog) error
	Update(ctx context.Context, log *model.OperationLog) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*model.OperationLog, error)
	List(ctx context.Context, opt *model.ListOption) ([]model.OperationLog, error)
	Count(ctx context.Context, opt *model.ListOption) (int64, error)
}
