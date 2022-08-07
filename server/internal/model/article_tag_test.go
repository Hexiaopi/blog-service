package model

import (
	"testing"

	"github.com/hexiaopi/blog-service/internal/config"
)

func TestArticleTag_Create(t *testing.T) {
	db, err := config.NewDBEngine(&config.TestDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	at := ArticleTag{ArticleID: 4, TagID: 2}
	if err := at.Create(db); err != nil {
		t.Fatal(err)
	}
}

func TestArticleTag_Update(t *testing.T) {
	db, err := config.NewDBEngine(&config.TestDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	at := ArticleTag{ArticleID: 4, TagID: 1}
	values := map[string]interface{}{"modified_by": "hexiaopi", "tag_id": 2}
	if err := at.Update(db, values); err != nil {
		t.Fatal(err)
	}
}

func TestArticleTag_Delete(t *testing.T) {
	db, err := config.NewDBEngine(&config.TestDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	at := ArticleTag{ArticleID: 4, TagID: 1}
	if err := at.Delete(db); err != nil {
		t.Fatal(err)
	}
}

func TestArticleTag_DeleteByArticle(t *testing.T) {
	db, err := config.NewDBEngine(&config.TestDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	at := ArticleTag{ArticleID: 3}
	if err := at.DeleteByArticle(db); err != nil {
		t.Fatal(err)
	}
}

func TestArticleTag_GetByArticle(t *testing.T) {
	db, err := config.NewDBEngine(&config.TestDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	at := ArticleTag{ArticleID: 4}
	result, err := at.GetByArticle(db)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
