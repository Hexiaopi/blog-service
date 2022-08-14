package model

import (
	"testing"

	"github.com/hexiaopi/blog-service/internal/config"
)

func TestAuth_Get(t *testing.T) {
	db, err := config.NewDBEngine(&config.TestDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	auth := Auth{AppKey: "abcd", AppSecret: "123456"}
	result, err := auth.Get(db)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.ID)
}
