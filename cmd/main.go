package main

import (
	config2 "github.com/kimxuanhong/user-manager-go/internal/config"
	"github.com/kimxuanhong/user-manager-go/internal/routes"
	"github.com/kimxuanhong/user-manager-go/internal/routes/route"
	"github.com/kimxuanhong/user-manager-go/pkg/app"
	"github.com/kimxuanhong/user-manager-go/pkg/dependencies"
)

func main() {
	cfg := config2.InitConfig()
	db := config2.InitDB(cfg)
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
	server.Start("127.0.0.1", "3001")
}
