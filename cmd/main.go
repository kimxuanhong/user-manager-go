package main

import (
	"github.com/kimxuanhong/user-manager-go/cmd/di"
	"github.com/kimxuanhong/user-manager-go/pkg/api/api"
	"github.com/kimxuanhong/user-manager-go/pkg/infra/config"
	"log"
)

func main() {
	app := di.NewDI()
	cfg := api.NewConfig()
	server := api.NewServer()
	database := config.NewDatasource(cfg)
	app.Dependencies(cfg)
	app.Dependencies(database)
	app.Dependencies(server)
	Application(server, app)
	if err := app.Graph(); err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	server.Start()
}
