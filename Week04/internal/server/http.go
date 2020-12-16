package server

import (
	"Week04/internal/dto"
	"Week04/internal/service"
	"Week04/pkg"
	"Week04/pkg/transport"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	*transport.Server
}

func NewServer(addr string, ser *service.ArticleService) *HttpServer {
	router := gin.New()

	//文章
	articleGroup := router.Group("/article")
	{
		articleGroup.GET("/:id", func(ctx *gin.Context) {
			validPram := &dto.ArticleQueryRequest{}
			if err := validPram.IsValid(ctx); err != nil {
				pkg.FailWithErr(ctx, err)
				return
			}
			aInfo, err := ser.Get(validPram.Id)
			if err != nil {
				pkg.FailWithErr(ctx, err)
				return
			}

			arResp := &dto.ArticleResponse{}
			pkg.Success(ctx, "", arResp.Output(aInfo))
		})

		articleGroup.POST("/create", func(ctx *gin.Context) {

		})
	}

	s := transport.NewHttpServer(addr, router)
	return &HttpServer{s}
}
