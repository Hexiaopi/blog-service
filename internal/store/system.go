package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/model"
)

type SystemConfigStore interface {
	Get(ctx context.Context, name string) (*model.Config, error)
}
