package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

const (
	ArticleFormat = "blog-service#article#total"
)

type ArticleCache struct {
	Client *redis.Client
}

func NewArticleCache(client *redis.Client) *ArticleCache {
	return &ArticleCache{Client: client}
}

func (c *ArticleCache) GetCount(ctx context.Context) (int64, error) {
	return c.Client.Get(ctx, ArticleFormat).Int64()
}

func (c *ArticleCache) SetCount(ctx context.Context, count int64) error {
	return c.Client.Set(ctx, ArticleFormat, count, 0).Err()
}
