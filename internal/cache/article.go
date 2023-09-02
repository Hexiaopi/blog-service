package cache

import (
	"context"
	"errors"
)

type ArticleCache interface {
	GetCount(ctx context.Context) (int64, error)
	SetCount(ctx context.Context, count int64) error
}

var ErrNotFound = errors.New("cache not found")
