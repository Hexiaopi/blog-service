package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/model"
)

type SysMenuStore interface {
	Create(ctx context.Context, sysMenu *model.SysMenu) error
	Update(ctx context.Context, sysMenu *model.SysMenu) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*model.SysMenu, error)
	List(ctx context.Context, opt *entity.ListOption) ([]model.SysMenu, error)
	Count(ctx context.Context, opt *entity.ListOption) (int64, error)
}
