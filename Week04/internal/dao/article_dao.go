package dao

import (
	"Week04/internal/ecode"
	"Week04/internal/model"

	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type ArticleDao struct {
	db *gorm.DB
}

func NewArticleDao(db *gorm.DB) *ArticleDao {
	return &ArticleDao{db: db}
}

//根据文章id 获取文章
func (a *ArticleDao) Get(id uint) (*model.Article, error) {
	articleInfo := &model.Article{}

	if err := a.db.First(articleInfo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(ecode.AirticleNotFound, "article not found:"+err.Error())
		}
		return nil, errors.Wrap(err, "get article failed")
	}
	return articleInfo, nil
}

//判断文章是否已存在
func (a *ArticleDao) Exist(id uint) (bool, error) {
	_, err := a.Get(id)
	if err != nil {
		if errors.Is(err, ecode.AirticleNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

//添加文章
func (a *ArticleDao) Create(article *model.Article) error {
	if err := a.db.Create(article).Error; err != nil {
		return errors.Wrap(err, "add article")
	}
	return nil
}

//更新文章
func (a *ArticleDao) Update(article *model.Article) error {
	if article.ID != 0 {
		found, err := a.Exist(article.ID)
		if err != nil {
			return errors.Wrap(err, "get article fail")
		}
		if !found {
			return errors.Wrap(ecode.AirticleNotFound, "article not found")
		}
	}
	if err := a.db.Save(article).Error; err != nil {
		return errors.Wrap(ecode.AirticleUpdateFail, "article update fail:"+err.Error())
	}
	return nil
}

//获取列表
func (a *ArticleDao) ListArticle(page, limit int) ([]*model.Article, error) {
	var articles []*model.Article
	if err := a.db.Find(articles).Limit((page - 1) * limit).Error; err != nil {
		return nil, errors.Wrap(ecode.AirticleListFail, "get article list fail:"+err.Error())
	}
	return articles, nil
}
