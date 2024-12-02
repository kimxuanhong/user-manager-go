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
	sql.InitPage[entity.User]().
		SetQuery(sql.GetUserByPartnerId).
		AndWhere("id = ?", request.RequestId).
		Fetch(ctx, func(obj *list.Array[entity.User], err error) {
			defer app.PanicHandler(func(obj any, err error) {
				whenDone(ctx, taskData, nil)
			})

			if err != nil {
				log.Printf("Query was error! %v\n", err)
				whenDone(ctx, taskData, nil)
				return
			}

			obj.ForEach(func(user entity.User) {
				log.Println(user.ID)
			})

			list.Map(obj, func(user entity.User) entity.User {
				return entity.User{
					ID: "test",
				}
			}).ForEach(func(user entity.User) {
				log.Println(user.ID)
			})

			taskData.Output = ctx.OK(obj.Slice())
			taskData.Input = nil
			whenDone(ctx, taskData, nil)
		})
}

func (r *GetUserByPartnerIdTask) GetName() string {
	return r.Name
}
