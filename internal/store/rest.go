package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/model"
)

type SysRestStore interface {
	Create(ctx context.Context, sysApi *model.SysRest) error
	Update(ctx context.Context, sysApi *model.SysRest) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*model.SysRest, error)
	List(ctx context.Context, opt *entity.ListOption) ([]model.SysRest, error)
	Count(ctx context.Context, opt *entity.ListOption) (int64, error)
}
