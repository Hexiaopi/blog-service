package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
)

type UserStore interface {
	Get(ctx context.Context, username string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) (int, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, opt *entity.ListOption) ([]entity.User, error)
	Count(ctx context.Context, opt *entity.ListOption) (int64, error)
}
