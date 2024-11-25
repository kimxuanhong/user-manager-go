package controller

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/pkg/api/api"
	"github.com/kimxuanhong/user-manager-go/pkg/infra/dao"
	"github.com/kimxuanhong/user-manager-go/pkg/infra/entity"
	"log"
	"sync"
)

type UserRoute interface {
	GetUserInfosByPartner(ctx *api.Context, whenDone api.Handler[any])
}

type userRoute struct {
	userDao dao.UserDao
}

var instanceUserRoute *userRoute
var userRouteOnce sync.Once

func NewUserRoute(userDao dao.UserDao) UserRoute {
	userRouteOnce.Do(func() {
		instanceUserRoute = &userRoute{userDao: userDao}
	})
	return instanceUserRoute
}

func (r *userRoute) GetUserInfosByPartner(ctx *api.Context, whenDone api.Handler[any]) {
	partnerId := ctx.Param("id")
	log.Println("Request " + partnerId)
	var req api.Request
	if err := ctx.Bind(&req); err != nil {
		whenDone(ctx.Bad(api.INVALID, err.Error()), nil)
		return
	}
	ctx.SetRequestId(req.RequestId)

	log.Println(req.RequestId)
	r.userDao.FindUserByPartnerId(ctx, partnerId, func(obj []entity.User, error error) {
		if error != nil || obj == nil {
			whenDone(nil, fmt.Errorf("bad request"))
			return
		}
		whenDone(ctx.OK(obj), nil)
	})
}
