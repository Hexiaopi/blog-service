package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/model"
)

type SysMenuStore interface {
	List(ctx context.Context, opt *entity.ListOption) ([]model.SysMenu, error)
}
