package model

import (
	"testing"

	"github.com/hexiaopi/blog-service/internal/config"
)

func TestArticle_Create(t *testing.T) {
	db, err := config.NewDBEngine(&config.TestDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	article := Article{
		Title:         "Docker实践",
		Desc:          "Docker学习开发实践",
		Content:       "docker...",
		CoverImageUrl: "",
		State:         1,
		CreatedBy:     "hexiaopi",
		ModifiedBy:    "",
		CreatedOn:     0,
		ModifiedOn:    0,
		DeletedOn:     0,
		IsDel:         0,
	}
	result, err := article.Create(db)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.ID)
}

func TestArticle_Get(t *testing.T) {
	db, err := config.NewDBEngine(&config.TestDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	article := Article{ID: 4, State: 1}
	result, err := article.Get(db)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestArticle_Update(t *testing.T) {
	db, err := config.NewDBEngine(&config.TestDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	values := make(map[string]interface{})
	article := Article{ID: 4}
	values["content"] = "docker从入门到精通"
	if err := article.Update(db, values); err != nil {
		t.Fatal(err)
	}
}

func TestArticle_Delete(t *testing.T) {
	db, err := config.NewDBEngine(&config.TestDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	article := Article{ID: 3}
	if err := article.Delete(db); err != nil {
		t.Fatal(err)
	}
}

func TestArticle_ListByTag(t *testing.T) {
	db, err := config.NewDBEngine(&config.TestDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	article := Article{State: 1}
	result, err := article.ListByTag(db, 1, 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range result {
		t.Log(*v)
	}
}

func TestArticle_CountByTag(t *testing.T) {
	db, err := config.NewDBEngine(&config.TestDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	article := Article{State: 1}
	count, err := article.CountByTag(db, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(count)
}
