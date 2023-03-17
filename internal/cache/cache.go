package cache

type Factory interface {
	Articles() ArticleCache
}
