package store

//go:generate mockgen -self_package=github.com/hexiaopi/blog-service/internal/store -destination mock_store.go -package store github.com/hexiaopi/blog-service/internal/store Factory,ArticleStore,TagStore,AuthStore

var client Factory

type Factory interface {
	Articles() ArticleStore
	Tags() TagStore
	Users() UserStore
	Systems() SystemConfigStore
	Close() error
}

func Client() Factory {
	return client
}

func SetClient(factory Factory) {
	client = factory
}
