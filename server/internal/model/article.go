package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	ID            uint32 `json:"id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	CreatedOn     uint32 `json:"created_on"`
	ModifiedOn    uint32 `json:"modified_on"`
	DeletedOn     uint32 `json:"deleted_on"`
	IsDel         uint8  `json:"is_del"`
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	db = db.Where("id = ? AND state = ? AND is_del= ?", a.ID, a.State, 0)
	err := db.Find(&article).Error
	if err != nil {
		return article, err
	}
	return article, nil
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("is_del= ?", 0).Update(values).Error; err != nil {
		return err
	}
	return nil
}

func (a Article) Delete(db *gorm.DB) error {
	if err := db.Where("is_del = ?", 0).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}

type ArticleEntity struct {
	ArticleID     uint32
	ArticleTitle  string
	ArticleDesc   string
	CoverImageUrl string
	Content       string
	TagID         uint32
	TagName       string
}

func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*ArticleEntity, error) {
	fields := []string{"ar.id AS article_id", "ar.title as article_title", "ar.desc AS article_desc", "ar.cover_image_url", "ar.content"}
	fields = append(fields, []string{"t.id AS tag_id", "t.name AS tag_name"}...)

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	rows, err := db.Select(fields).Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id = ar.id").
		Where("ar.state = ? AND ar.is_del = ?", a.State, 0).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*ArticleEntity
	for rows.Next() {
		r := &ArticleEntity{}
		if err := rows.Scan(&r.ArticleID, &r.ArticleTitle, &r.ArticleDesc, &r.CoverImageUrl, &r.Content, &r.TagID, &r.TagName); err != nil {
			return nil, err
		}
		articles = append(articles, r)
	}
	return articles, nil
}

func (a Article) Count(db *gorm.DB) (int, error) {
	var count int
	err := db.Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id = ar.id").
		Where("ar.state = ? AND ar.is_del = ?", a.State, 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (a Article) ListByTag(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*ArticleEntity, error) {
	fields := []string{"ar.id AS article_id", "ar.title as article_title", "ar.desc AS article_desc", "ar.cover_image_url", "ar.content"}
	fields = append(fields, []string{"t.id AS tag_id", "t.name AS tag_name"}...)

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	rows, err := db.Select(fields).Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id = ar.id").
		Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*ArticleEntity
	for rows.Next() {
		r := &ArticleEntity{}
		if err := rows.Scan(&r.ArticleID, &r.ArticleTitle, &r.ArticleDesc, &r.CoverImageUrl, &r.Content, &r.TagID, &r.TagName); err != nil {
			return nil, err
		}
		articles = append(articles, r)
	}
	return articles, nil
}

func (a Article) CountByTag(db *gorm.DB, tagID uint32) (int, error) {
	var count int
	err := db.Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id = ar.id").
		Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
