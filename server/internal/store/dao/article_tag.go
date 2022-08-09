package dao

// func (d *Dao) GetArticleTagByAID(articleID uint32) (model.ArticleTag, error) {
// 	articleTag := model.ArticleTag{ArticleID: articleID}
// 	return articleTag.GetByArticle(d.engine)
// }

// func (d *Dao) CreateArticleTag(articleID, tagID uint32, createBy string) error {
// 	articleTag := model.ArticleTag{
// 		TagID:     tagID,
// 		ArticleID: articleID,
// 		CreatedBy: createBy,
// 	}
// 	return articleTag.Create(d.engine)
// }

// func (d *Dao) UpdateArticleTag(articleID, tagID uint32, modifiedBy string) error {
// 	articleTag := model.ArticleTag{ArticleID: articleID, TagID: tagID}
// 	values := map[string]interface{}{
// 		"article_id":  articleID,
// 		"tag_id":      tagID,
// 		"modified_by": modifiedBy,
// 	}
// 	return articleTag.Update(d.engine, values)
// }

// func (d *Dao) DeleteArticleTag(articleID uint32) error {
// 	articleTag := model.ArticleTag{ArticleID: articleID}
// 	return articleTag.DeleteByArticle(d.engine)
// }
