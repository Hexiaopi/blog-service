package store

import "context"

//go:generate mockgen -self_package=github.com/hexiaopi/blog-service/internal/store -destination mock_store.go -package store github.com/hexiaopi/blog-service/internal/store Factory,ArticleStore,TagStore,UserStore,RoleStore,SystemConfigStore,ResourceStore,OperationStore,UserRoleStore

var client Factory

type Factory interface {
	Articles() ArticleStore
	Tags() TagStore
	Users() UserStore
	Roles() RoleStore
	Systems() SystemConfigStore
	Resources() ResourceStore
	Operations() OperationStore
	UserRole() UserRoleStore
	Close() error
	Tx(ctx context.Context, f TxFunc) error
}

func Client() Factory {
	return client
}

func SetClient(factory Factory) {
	client = factory
}

type TxFunc = func(ctx context.Context, factory Factory) error
