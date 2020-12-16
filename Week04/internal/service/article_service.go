package service

import (
	"Week04/internal/dao"
	"Week04/internal/model"
)

type ArticleService struct {
	dao *dao.ArticleDao
}

func NewArticleService(dao *dao.ArticleDao) *ArticleService {
	return &ArticleService{dao: dao}
}

func (a *ArticleService) Get(id uint) (*model.Article, error) {
	return a.dao.Get(id)
}

func (a *ArticleService) Create(article *model.Article) error {
	return a.dao.Create(article)
}

func (a *ArticleService) Update(article *model.Article) error {
	return a.dao.Update(article)
}

func (a *ArticleService) ListArticle(page, limit int) ([]*model.Article, error) {
	return a.dao.ListArticle(page, limit)
}
