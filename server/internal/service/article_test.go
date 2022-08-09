package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestListArticle(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockObj := store.NewMockArticleStore(mockCtrl)

	articles := []*entity.Article{{ID: 1}, {ID: 2}}

	mockObj.EXPECT().List(gomock.Any(), gomock.Any()).Return(articles, int64(2), nil)

	ret, total, err := mockObj.List(context.Background(), &entity.ListOption{})
	assert.Equal(t, err, nil)
	assert.Equal(t, ret, articles)
	assert.Equal(t, total, int64(2))
}
