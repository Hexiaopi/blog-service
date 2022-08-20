package model

type ArticleTag struct {
	ID        int `json:"id"`
	TagID     int `json:"tag_id"`
	ArticleID int `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
