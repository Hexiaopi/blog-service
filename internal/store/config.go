package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
)

type SystemConfigStore interface {
	Get(ctx context.Context, name string) (*entity.Config, error)
}
