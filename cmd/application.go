package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kimxuanhong/user-manager-go/cmd/di"
	"github.com/kimxuanhong/user-manager-go/pkg/api/api"
	"github.com/kimxuanhong/user-manager-go/pkg/api/controller"
	"github.com/kimxuanhong/user-manager-go/pkg/core/service"
	"github.com/kimxuanhong/user-manager-go/pkg/infra/repository"
)

func Application(server api.Server, newDI di.DI) {
	//Dao
	newDI.Dependencies(repository.NewUserRepository())

	//Service layer
	newDI.Dependencies(service.NewUserService())

	//API layer
	userHandler := controller.NewUserController()
	newDI.Dependencies(userHandler)
	server.Router("/api/v1/ctx/users", func(router *gin.RouterGroup) {
		router.GET("/", api.Handler(userHandler.GetUserInfo))
		router.PUT("/:id", api.Handler(userHandler.UpdateUserStatus))
	})
}
