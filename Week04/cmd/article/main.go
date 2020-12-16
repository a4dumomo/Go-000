package main

import (
	"Week04/pkg"
	"context"
	"log"
)

func main() {

	path := "./../../configs/config.yaml"
	conf := pkg.NewConfig(path)
	pkg.NewDb(conf)

	httpSrv := InitNewAir(conf.HttpServer.Addr, pkg.Db)

	app := pkg.New()
	app.Append(pkg.Hook{
		OnStart: func(ctx context.Context) error {
			return httpSrv.Starrt()
		},
		OnStop: func(ctx context.Context) error {
			return httpSrv.Stop(ctx)
		},
	})
	if err := app.Run(); err != nil {
		log.Fatal("service ecode:", err)
	}
}
