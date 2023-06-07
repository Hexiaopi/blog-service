package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/model"
)

type OperationStore interface {
	Create(ctx context.Context, log *model.SystemOperationLog) error
	Update(ctx context.Context, log *model.SystemOperationLog) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*model.SystemOperationLog, error)
	List(ctx context.Context, opt *model.ListOption) ([]model.SystemOperationLog, error)
	Count(ctx context.Context, opt *model.ListOption) (int64, error)
}
