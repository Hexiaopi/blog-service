package store

//go:generate mockgen -self_package=github.com/hexiaopi/blog-service/internal/store -destination mock_store.go -package store github.com/hexiaopi/blog-service/internal/store Factory,ArticleStore,TagStore,UserStore,RoleStore,SystemConfigStore,ResourceStore,OperationStore

var client Factory

type Factory interface {
	Articles() ArticleStore
	Tags() TagStore
	Users() UserStore
	Roles() RoleStore
	Systems() SystemConfigStore
	Resources() ResourceStore
	Operations() OperationStore
	Close() error
}

func Client() Factory {
	return client
}

func SetClient(factory Factory) {
	client = factory
}
