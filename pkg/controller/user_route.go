package controller

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/internal/infra/dao"
	"github.com/kimxuanhong/user-manager-go/pkg/api"
	"github.com/kimxuanhong/user-manager-go/pkg/dto"
	"github.com/kimxuanhong/user-manager-go/pkg/entity"
	"sync"
)

type getUserByIdRoute struct {
	userDao dao.UserDao
}

var instanceUserRoute *getUserByIdRoute
var userRouteOnce sync.Once

func GetUserByIdRoute(userDao dao.UserDao) api.Route {
	userRouteOnce.Do(func() {
		instanceUserRoute = &getUserByIdRoute{userDao: userDao}
	})
	return instanceUserRoute
}

func (r *getUserByIdRoute) RouteHandler(ctx *api.Context, whenDone api.Handler[any]) {
	partnerId := ctx.Param("id")
	var req dto.Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		whenDone(ctx.Bad(dto.INVALID, err.Error()), nil)
		return
	}
	ctx.SetRequestId(req.RequestId)
	r.userDao.FindUserByPartnerId(ctx, partnerId, func(obj []entity.User, error error) {
		if error != nil || obj == nil {
			whenDone(ctx.Bad(dto.INVALID, fmt.Errorf("bad request").Error()), nil)
			return
		}
		whenDone(ctx.OK(obj), nil)
	})
}
