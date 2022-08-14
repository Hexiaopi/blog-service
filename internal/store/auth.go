package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/model"
)

type AuthStore interface {
	Get(ctx context.Context, key, secret string) (*model.Auth, error)
}
