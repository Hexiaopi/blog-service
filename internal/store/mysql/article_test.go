package mysql

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/hexiaopi/blog-service/internal/model"
)

func TestArticleDao_Create(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	article := model.Article{
		Name:       "test",
		Desc:       "test",
		Content:    "test content",
		State:      model.STATE_OPEN,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Operator:   "admin",
	}
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `blog_article` (`name`,`desc`,`content`,`state`,`create_time`,`update_time`,`operator`) VALUES (?,?,?,?,?,?,?)").
		WithArgs(article.Name, article.Desc, article.Content, article.State, sqlmock.AnyArg(), sqlmock.AnyArg(), article.Operator).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	gdb, err := gorm.Open(mysql.New(
		mysql.Config{DriverName: "mysql", Conn: db, SkipInitializeWithVersion: true}),
		&gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	dao := NewArticleDao(gdb)
	if err := dao.Create(context.Background(), &article); err != nil {
		t.Fatal(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestArticleDao_Get(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	mock.ExpectQuery("SELECT * FROM `blog_article` WHERE `blog_article`.`id` = ? ORDER BY `blog_article`.`id` LIMIT 1").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "desc", "content", "state", "create_time", "update_time", "operator"}).
			AddRow("1", "test", "test", "test content", 1, time.Now(), time.Now(), "admin"))
	mock.ExpectQuery("SELECT * FROM `blog_article_tag` WHERE `blog_article_tag`.`article_id` = ?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "article_id", "tag_id"}).
			AddRow("1", "1", "1"))
	mock.ExpectQuery("SELECT * FROM `blog_tag` WHERE `blog_tag`.`id` = ?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "desc", "state", "create_time", "update_time", "operator"}).
			AddRow("1", "test", "test", 1, time.Now(), time.Now(), "admin"))

	gdb, err := gorm.Open(mysql.New(
		mysql.Config{DriverName: "mysql", Conn: db, SkipInitializeWithVersion: true}),
		&gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	dao := NewArticleDao(gdb)
	article, err := dao.Get(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(article)
}

func TestArticleDao_Update(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	article := model.Article{
		ID:      1,
		Name:    "test",
		Desc:    "test",
		Content: "test content",
	}
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `blog_article` SET `name`=?,`desc`=?,`content`=?,`update_time`=? WHERE `id` = ?").
		WithArgs(article.Name, article.Desc, article.Content, sqlmock.AnyArg(), article.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	gdb, err := gorm.Open(mysql.New(
		mysql.Config{DriverName: "mysql", Conn: db, SkipInitializeWithVersion: true}),
		&gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	dao := NewArticleDao(gdb)
	if err := dao.Update(context.Background(), &article); err != nil {
		t.Fatal(err)
	}
}

func TestArticleDao_Delete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `blog_article` WHERE `blog_article`.`id` = ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	gdb, err := gorm.Open(mysql.New(
		mysql.Config{DriverName: "mysql", Conn: db, SkipInitializeWithVersion: true}),
		&gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	dao := NewArticleDao(gdb)
	if err := dao.Delete(context.Background(), 1); err != nil {
		t.Fatal(err)
	}
}

func TestArticleDao_Count(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	opt := model.ListOption{
		State: 10,
	}

	mock.ExpectQuery("SELECT count(*) FROM `blog_article` WHERE state = ?").
		WithArgs(opt.State).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).
			AddRow(7))

	gdb, err := gorm.Open(mysql.New(
		mysql.Config{DriverName: "mysql", Conn: db, SkipInitializeWithVersion: true}),
		&gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	dao := NewArticleDao(gdb)
	count, err := dao.Count(context.Background(), &opt)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(count)
}
