package store

import "context"

type RoleMenuStore interface {
	Create(ctx context.Context, roleId, menuId int) error
	DeleteByRole(ctx context.Context, roleId int) error
	ListByRole(ctx context.Context, roleId int) ([]int, error)
}
