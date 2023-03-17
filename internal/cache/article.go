package cache

import "context"

type ArticleCache interface {
	GetCount(ctx context.Context) (int64, error)
	SetCount(ctx context.Context, count int64) error
}
