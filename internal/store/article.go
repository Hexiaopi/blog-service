package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/model"
)

type ArticleStore interface {
	Create(ctx context.Context, article *model.Article) error
	Update(ctx context.Context, article *model.Article) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*model.Article, error)
	List(ctx context.Context, opt *model.ListOption) ([]model.Article, int64, error)
}
