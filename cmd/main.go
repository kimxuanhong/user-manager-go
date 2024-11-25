package main

import (
	"github.com/kimxuanhong/user-manager-go/pkg/api/api"
	"github.com/kimxuanhong/user-manager-go/pkg/api/config"
	"github.com/kimxuanhong/user-manager-go/pkg/api/controller"
	"github.com/kimxuanhong/user-manager-go/pkg/infra/dao"
)

func main() {
	cfg := config.NewConfig()
	db := config.NewDatasource(cfg)
	deps := &config.Dependencies{
		Config: cfg,
		DB:     db,
	}
	userDao := dao.NewUserDao()
	userHandler := controller.NewUserRoute(userDao)
	server := api.NewHttpServer(deps)
	server.Post("/partner/:id", userHandler.GetUserInfosByPartner)
	server.Start("127.0.0.1", "3001")
}
