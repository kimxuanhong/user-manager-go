package route

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/internal/core/workflow"
	"github.com/kimxuanhong/user-manager-go/internal/dto"
	"github.com/kimxuanhong/user-manager-go/internal/infra/dao"
	"github.com/kimxuanhong/user-manager-go/internal/infra/entity"
	"github.com/kimxuanhong/user-manager-go/pkg/app"
	"github.com/kimxuanhong/user-manager-go/pkg/task"
	"sync"
)

type PartnerRoute interface {
	GetUserByPartnerId(ctx *app.Context, whenDone app.Handler[any])
	GetAllUser(ctx *app.Context, whenDone app.Handler[any])
}

type partnerRoute struct {
	userDao dao.UserDao
}

var instancePartnerRoute *partnerRoute
var partnerRouteOnce sync.Once

func NewPartnerRoute(userDao dao.UserDao) PartnerRoute {
	partnerRouteOnce.Do(func() {
		instancePartnerRoute = &partnerRoute{userDao: userDao}
	})
	return instancePartnerRoute
}

func (r *partnerRoute) GetUserByPartnerId(ctx *app.Context, whenDone app.Handler[any]) {
	partnerId := ctx.Param("id")
	var req dto.Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		whenDone(ctx.Bad(dto.INVALID, err.Error()), nil)
		return
	}
	ctx.SetRequestId(req.RequestId)

	r.userDao.FindUserByPartnerId(ctx, partnerId, func(obj []entity.User, err error) {
		if err != nil || obj == nil {
			whenDone(ctx.Bad(dto.INVALID, fmt.Errorf("bad request").Error()), nil)
			return
		}
		wf := workflow.NewMyWorkFlow()
		data := &obj[0]
		wf.Run(ctx, &task.Data{
			Input:  data,
			Output: data,
		}, func(ctx *app.Context, taskData *task.Data, err error) {
			whenDone(ctx.OK(obj), nil)
		})

	})
}

func (r *partnerRoute) GetAllUser(ctx *app.Context, whenDone app.Handler[any]) {
	var req dto.Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		whenDone(ctx.Bad(dto.INVALID, err.Error()), nil)
		return
	}
	ctx.SetRequestId(req.RequestId)
	r.userDao.FindAllUser(ctx, func(obj []entity.User, error error) {
		if error != nil || obj == nil {
			whenDone(ctx.Bad(dto.INVALID, fmt.Errorf("bad request").Error()), nil)
			return
		}
		whenDone(ctx.OK(obj), nil)
	})
}
