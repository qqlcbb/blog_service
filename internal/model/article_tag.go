package model

import "github.com/jinzhu/gorm"

type ArticleTag struct {
	*Model
	TagID uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

func (a ArticleTag) GetByID(db *gorm.DB) (ArticleTag, error) {
	var articleTag ArticleTag
	err := db.Where("article_id = ? and is_del = 0", a.ArticleID).First(&articleTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return articleTag, err
	}
	return articleTag, nil
}

func (a ArticleTag) ListByTID(db *gorm.DB) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	err := db.Where("tag_id = ? and is_del = 0", a.TagID).Find(&articleTags).Error
	if err != nil {
		return nil, err
	}
	return articleTags, nil
}

func (a ArticleTag) ListByAIDS(db *gorm.DB, articleIDs []uint32) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	err := db.Where("article_id in (?) and is_del = 0", articleIDs).Find(&articleTags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articleTags, nil
}

func (a ArticleTag) Create (db *gorm.DB) error {
	err := db.Create(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func (a ArticleTag) UpdateOne(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("article_id = ? and is_del = 0", a.ArticleID).Limit(1).Update(values).Error; err != nil {
		return err
	}
	return nil
}

func (a ArticleTag) Delete (db *gorm.DB) error {
	if err := db.Where("id = ? and is_del = 0", a.Model.ID).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}

func (a ArticleTag) DeleteOne(db *gorm.DB) error {
	if err := db.Where("article_id = ? and is_del = 0", a.ArticleID).Delete(&a).Limit(1).Error; err != nil {
		return err
	}
	return nil
}
