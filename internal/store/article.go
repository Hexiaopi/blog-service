package store

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
)

type ArticleStore interface {
	Create(ctx context.Context, article *entity.Article) error
	Update(ctx context.Context, article *entity.Article) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*entity.Article, error)
	List(ctx context.Context, opt *entity.ListOption) ([]entity.Article, int64, error)
}
