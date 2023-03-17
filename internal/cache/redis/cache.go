package redis

import (
	"github.com/go-redis/redis/v8"

	"github.com/hexiaopi/blog-service/internal/cache"
)

type CacheDao struct {
	client *redis.Client
}

func NewCache(client *redis.Client) *CacheDao {
	return &CacheDao{client: client}
}

func (dao *CacheDao) Articles() cache.ArticleCache {
	return NewArticleCache(dao.client)
}
