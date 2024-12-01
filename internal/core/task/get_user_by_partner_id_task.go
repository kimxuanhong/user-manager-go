package task

import (
	"github.com/kimxuanhong/user-manager-go/internal/dto"
	"github.com/kimxuanhong/user-manager-go/internal/infra/entity"
	"github.com/kimxuanhong/user-manager-go/internal/infra/sql"
	"github.com/kimxuanhong/user-manager-go/pkg/app"
	"github.com/kimxuanhong/user-manager-go/pkg/task"
	"log"
)

type GetUserByPartnerIdTask struct {
	Name string
}

func NewGetUserByPartnerIdTask() task.Task {
	return &GetUserByPartnerIdTask{
		Name: "GetConfigTask",
	}
}

func (r *GetUserByPartnerIdTask) Execute(ctx *app.Context, taskData *task.Data, whenDone task.Handler) {
	request := taskData.Input.(*dto.Request)
	sql.QueryWithParams(ctx, sql.Params{Query: sql.GetUserByPartnerId, Values: []interface{}{request.RequestId}}, func(users []entity.User, err error) {
		task.SafeCallback(ctx, whenDone, func() {
			if err != nil {
				log.Printf("Query was error! %v\n", err)
				whenDone(ctx, taskData, err)
				return
			}
			taskData.Output = ctx.OK(&users[0])
			whenDone(ctx, taskData, nil)
		})
	})
}

func (r *GetUserByPartnerIdTask) GetName() string {
	return r.Name
}
