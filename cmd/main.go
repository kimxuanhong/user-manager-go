package main

import (
	"github.com/kimxuanhong/user-manager-go/internal/config"
	"github.com/kimxuanhong/user-manager-go/internal/routes"
	"github.com/kimxuanhong/user-manager-go/internal/routes/route"
	"github.com/kimxuanhong/user-manager-go/pkg/app"
	"github.com/kimxuanhong/user-manager-go/pkg/dependencies"
)

func main() {
	cfg := config.InitConfig()
	db := config.InitDB(cfg)
	defer config.CloseDB(db)
	partnerRoutes := route.NewPartnerRoute()

	app.LogWorker()
	defer app.OnStopServer()
	server := app.NewHttpServer(&dependencies.Dependency{
		Db: db,
	})
	server.Middleware(app.RecoveryMiddleware())
	server.Middleware(app.LogRequestMiddleware())
	server.Routes(routes.PartnerRoutes(partnerRoutes))
	server.Middleware(app.LogResponseMiddleware())
	server.Start(cfg.Server.Host, cfg.Server.Port)
}
