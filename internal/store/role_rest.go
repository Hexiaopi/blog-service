package store

import "context"

type RoleRestStore interface {
	Create(ctx context.Context, roleId, restId int) error
	DeleteByRole(ctx context.Context, roleId int) error
	ListByRole(ctx context.Context, roleId int) ([]int, error)
}
