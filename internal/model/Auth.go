package model

import "github.com/jinzhu/gorm"

type Auth struct {
	*Model
	AppKey string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (Auth) TableName() string {
	return "blog_auth"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? and app_secret = ? and is_del = 0", a.AppKey, a.AppSecret)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}