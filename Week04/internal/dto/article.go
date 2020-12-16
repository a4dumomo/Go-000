package dto

import (
	"Week04/internal/ecode"
	"Week04/internal/model"
	"time"

	"github.com/gin-gonic/gin"
)

type ArticleQueryRequest struct {
	Id uint `uri:"id" form:"id" json:"id" binding:"required"`
}

func (ar *ArticleQueryRequest) IsValid(ctx *gin.Context) error {
	if err := ctx.ShouldBindUri(ar); err != nil {
		return ecode.AirtilceIdIsError
	}
	return nil
}

//输出转换
type ArticleResponse struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UserId      int    `json:"user_id"`      // 用户
	Tag         int    `json:"tag"`          //分类
	Content     string `json:"content"`      //内容
	ClickNumber int    `json:"click_number"` // 点击量
	ZanNumber   int    `json:"zan_number"`   //点赞数
}

func (arResp *ArticleResponse) Output(article *model.Article) *ArticleResponse {
	arResp.ID = article.ID
	arResp.CreatedAt = article.CreatedAt
	arResp.UserId = article.UserId
	arResp.Tag = article.Tag
	arResp.Content = article.Content
	arResp.ClickNumber = article.ClickNumber
	arResp.ZanNumber = article.ZanNumber
	return arResp
}

//输入请求
type ArticleRequest struct {
	Tag     int    `form:"tag" json:"tag" binding:"required"`         //分类
	Content string `form:"content" json:"content" binding:"required"` //内容
}

func (ar *ArticleRequest) IsValid(ctx gin.Context) error {
	if err := ctx.ShouldBindJSON(ar); err != nil {
		return ecode.AirtilceIdIsError
	}
	return nil
}
