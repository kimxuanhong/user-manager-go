package controller

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/internal/infra/dao"
	"github.com/kimxuanhong/user-manager-go/pkg/api"
	"github.com/kimxuanhong/user-manager-go/pkg/dto"
	"github.com/kimxuanhong/user-manager-go/pkg/entity"
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
	var req dto.Request
	if err := ctx.Bind(&req); err != nil {
		whenDone(ctx.Bad(dto.INVALID, err.Error()), nil)
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
