package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/model"
)

type UserStore interface {
	Get(ctx context.Context, username string) (*model.User, error)
}
