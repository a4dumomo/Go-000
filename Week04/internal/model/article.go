package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	UserId      int    // 用户
	Tag         int    //分类
	Content     string //内容
	ClickNumber int    // 点击量
	ZanNumber   int    //点赞数
}

func (a *Article) TableName() string {
	return "article"
}
