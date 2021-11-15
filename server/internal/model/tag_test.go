package model

import (
	"blog-service/internal/config"
	"testing"
)

func TestTag_Create(t *testing.T) {
	db, err := config.NewDBEngine(&config.TestDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	tag := Tag{
		Name:      "Docker",
		State:     1,
		CreatedBy: "hexiaopi",
	}
	result, err := tag.Create(db)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.ID)
}

func TestTag_Get(t *testing.T) {
	db, err := config.NewDBEngine(&config.TestDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	tag := Tag{ID: 2, State: 1}
	result, err := tag.Get(db)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestTag_Update(t *testing.T) {
	db, err := config.NewDBEngine(&config.TestDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	tag := Tag{ID: 2}
	values := make(map[string]interface{})
	values["modified_by"] = "hexiaopi"
	if err := tag.Update(db, values); err != nil {
		t.Fatal(err)
	}
}

func TestTag_Delete(t *testing.T) {
	db, err := config.NewDBEngine(&config.TestDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	tag := Tag{ID: 3}
	if err := tag.Delete(db); err != nil {
		t.Fatal(err)
	}
}

func TestTag_Count(t *testing.T) {
	db, err := config.NewDBEngine(&config.TestDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	tag := Tag{State: 1}
	count, err := tag.Count(db)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(count)
}

func TestTag_List(t *testing.T) {
	db, err := config.NewDBEngine(&config.TestDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	tag := Tag{State: 1}
	tags, err := tag.List(db, 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	for _, tag := range tags {
		t.Log(tag)
	}

}
