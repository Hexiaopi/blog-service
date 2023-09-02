package redis

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"

	"github.com/hexiaopi/blog-service/internal/cache"
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
	count, err := c.Client.Get(ctx, ArticleFormat).Int64()
	if errors.Is(err, redis.Nil) {
		return 0, cache.ErrNotFound
	}
	return count, err
}

func (c *ArticleCache) SetCount(ctx context.Context, count int64) error {
	return c.Client.Set(ctx, ArticleFormat, count, 0).Err()
}
