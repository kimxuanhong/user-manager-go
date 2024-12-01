package route

import (
	"github.com/kimxuanhong/user-manager-go/internal/core/workflow"
	"github.com/kimxuanhong/user-manager-go/internal/dto"
	"github.com/kimxuanhong/user-manager-go/pkg/app"
	"github.com/kimxuanhong/user-manager-go/pkg/task"
	"sync"
)

type PartnerRoute interface {
	GetUserByPartnerId(ctx *app.Context, whenDone app.Handler[any])
	GetAllUser(ctx *app.Context, whenDone app.Handler[any])
}

type partnerRoute struct {
}

var instancePartnerRoute *partnerRoute
var partnerRouteOnce sync.Once

func NewPartnerRoute() PartnerRoute {
	partnerRouteOnce.Do(func() {
		instancePartnerRoute = &partnerRoute{}
	})
	return instancePartnerRoute
}

func (r *partnerRoute) GetUserByPartnerId(ctx *app.Context, whenDone app.Handler[any]) {
	var req dto.Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		whenDone(ctx.Bad(dto.INVALID, err.Error()), nil)
		return
	}
	ctx.SetRequestId(req.RequestId)

	wf := workflow.NewMyWorkFlow()
	wf.Run(ctx, &task.Data{
		Input:  &req,
		Output: &dto.Response{},
	}, func(ctx *app.Context, taskData *task.Data, err error) {
		if err != nil {
			whenDone(ctx.Bad(dto.INVALID, err.Error()), nil)
			return
		}
		whenDone(taskData.Output, nil)
	})

}

func (r *partnerRoute) GetAllUser(ctx *app.Context, whenDone app.Handler[any]) {
	var req dto.Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		whenDone(ctx.Bad(dto.INVALID, err.Error()), nil)
		return
	}
	ctx.SetRequestId(req.RequestId)
	whenDone(ctx.OK(req), nil)

}
