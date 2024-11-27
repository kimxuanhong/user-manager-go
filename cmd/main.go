package main

import (
	"github.com/kimxuanhong/user-manager-go/internal/infra/dao"
	"github.com/kimxuanhong/user-manager-go/pkg/api"
	"github.com/kimxuanhong/user-manager-go/pkg/config"
	"github.com/kimxuanhong/user-manager-go/pkg/controller"
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
