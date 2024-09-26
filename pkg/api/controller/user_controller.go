package controller

import (
	"github.com/kimxuanhong/user-manager-go/pkg/api/api"
	"github.com/kimxuanhong/user-manager-go/pkg/core/service"
	"sync"
)

type UserController interface {
	GetUserInfo(g *api.Context)
	UpdateUserStatus(g *api.Context)
}

type userController struct {
	UserService service.UserService `inject:""`
}

var instanceUserController *userController
var userControllerOnce sync.Once

func NewUserController() UserController {
	userControllerOnce.Do(func() {
		instanceUserController = &userController{}
	})
	return instanceUserController
}

func (r *userController) GetUserInfo(g *api.Context) {
	// Implement logic here
}

func (r *userController) UpdateUserStatus(g *api.Context) {
	// Implement logic here
}
