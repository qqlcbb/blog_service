package model

import "github.com/jinzhu/gorm"

type Article struct {
	*Model
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State uint8 `json:"state"`
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

func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Update(values).Where("id = ? and is_del = 0", a.ID).Error; err != nil {
		return err
	}
	return nil
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	err := db.Where("id = ? and state = ? and is_del = 0", a.ID, a.State).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}

	return article, nil
}

func (a Article) Delete (db *gorm.DB) error {
	if err := db.Where("id = ? and is_del = 0", a.ID).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}

type ArticleRow struct {
	ArticleID uint32
	TagID uint32
	TagName string
	ArticleTitle string
	ArticleDesc string
	CoverImageUrl string
	Content string
}

func (a Article) ListByTagID(db *gorm.DB, TagID uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {
	fields := []string{"ar.id as article_id", "ar.title as article_title", "ar.desc as article_desc", "ar.cover_image_url", "ar.content"}
	fields = append(fields, []string{"t.id as tag_id", "t.name as tag_name"}...)

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	rows, err := db.Select(fields).Table(ArticleTag{}.TableName() + " as at").
		Joins("left join `"+Tag{}.TableName()+"` as t on at.tag_id = t.id").
		Joins("left join `"+Article{}.TableName()+"` as ar on at.article_id = ar.id").
		Where("at.tag_id = ? and ar.state = ? and ar.is_del = 0", TagID, a.State).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*ArticleRow
	for rows.Next() {
		r := &ArticleRow{}
		if err := rows.Scan(&r.ArticleID, &r.ArticleTitle, &r.ArticleDesc, &r.CoverImageUrl, &r.Content, &r.TagID, &r.TagName); err != nil {
			return nil, err
		}
		articles = append(articles, r)
	}
	return articles, nil
}

func (a Article) CountByTagID(db *gorm.DB, TagID uint32) (int, error) {
	var count int
	err := db.Table(ArticleTag{}.TableName()+" as at").
		Joins("left join `"+Tag{}.TableName()+"` as t on at.tag_id = t.id").
		Joins("left join `"+Article{}.TableName()+"` as ar on at.article_id = ar.id").
		Where("at.tag_id = ? and ar.state = ? and ar.is_del = 0", TagID, a.State).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}





