//+build wireinject

package main

import (
	"Week04/internal/dao"
	"Week04/internal/server"
	"Week04/internal/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitNewAir(addr string, db *gorm.DB) *server.HttpServer {
	wire.Build(dao.NewArticleDao, service.NewArticleService, server.NewServer)
	return &server.HttpServer{}
}
