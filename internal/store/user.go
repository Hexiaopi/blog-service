package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/model"
)

type UserStore interface {
	Get(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, opt *model.ListOption) ([]model.User, error)
	Count(ctx context.Context, opt *model.ListOption) (int64, error)
}
