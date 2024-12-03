package task

import (
	"github.com/kimxuanhong/user-manager-go/internal/dto"
	"github.com/kimxuanhong/user-manager-go/internal/infra/entity"
	"github.com/kimxuanhong/user-manager-go/internal/infra/sql"
	"github.com/kimxuanhong/user-manager-go/pkg/app"
	"github.com/kimxuanhong/user-manager-go/pkg/task"
	"github.com/kimxuanhong/user-manager-go/pkg/utils/list"
	"log"
)

type GetUserByPartnerIdTask struct {
	Name string
}

func NewGetUserByPartnerIdTask() task.Task {
	return &GetUserByPartnerIdTask{
		Name: "GetUserByPartnerIdTask",
	}
}

func (r *GetUserByPartnerIdTask) Execute(ctx *app.Context, taskData *task.Data, whenDone task.Handler) {
	request := taskData.Input.(*dto.Request)
	sql.InitPage[*entity.User]().
		SetQuery(sql.GetUserByPartnerId).
		SetPageNumber(request.PageNumber).
		SetPageSize(request.PageSize).
		AndWhere("id = ?", request.RequestId).
		Fetch(ctx, app.SafeCallback(func(obj *sql.Page[*entity.User], err error) {
			if err != nil {
				log.Printf("Query was error! %v\n", err)
				whenDone(ctx, taskData, nil)
				return
			}

			obj.Data.ForEach(func(user *entity.User) {
				log.Println(user.ID)
			})

			list.Map(obj.Data, func(user *entity.User) *entity.User {
				return &entity.User{
					ID: "test",
				}
			}).ForEach(func(user *entity.User) {
				log.Println(user.ID)
			})

			taskData.Output = ctx.OK(obj)
			taskData.Input = nil
			whenDone(ctx, taskData, nil)
		}))
}

func (r *GetUserByPartnerIdTask) GetName() string {
	return r.Name
}
