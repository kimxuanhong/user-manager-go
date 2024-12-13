package task

import (
	"encoding/json"
	"github.com/kimxuanhong/user-manager-go/internal/dto"
	"github.com/kimxuanhong/user-manager-go/internal/infra/entity"
	"github.com/kimxuanhong/user-manager-go/internal/infra/sql"
	"github.com/kimxuanhong/user-manager-go/pkg/app"
	"github.com/kimxuanhong/user-manager-go/pkg/utils/hashmap"
	"github.com/kimxuanhong/user-manager-go/pkg/utils/list"
	"log"
)

type GetUserByPartnerIdTask struct {
	Name string
}

func NewGetUserByPartnerIdTask() Task {
	return &GetUserByPartnerIdTask{
		Name: "GetUserByPartnerIdTask",
	}
}

func (r *GetUserByPartnerIdTask) Execute(ctx *app.Context, taskData *Data, whenDone Handler) {
	request := taskData.Request.(*dto.Request)
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

			for item := range obj.Data.Iter() {
				log.Println("Item ...." + item.ID)
			}

			list.Map(obj.Data, func(user *entity.User) *entity.User {
				return &entity.User{
					ID: "test",
				}
			}).ForEach(func(user *entity.User) {
				log.Println(user.ID)
			})

			newMap := hashmap.NewMap[string, string]()

			newMap.Put("abc", "abc")

			log.Printf("map %s", newMap)
			jsonInput := `{
    "request_id": "78c83478-5e15-4720-9acb-b70ab32f011b",
    "PartnerId": "",
    "UserName": "",
    "FirstName": "",
    "LastName": "",
    "Email": "",
    "Status": "",
    "request_time": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z"
}`
			for imap := range newMap.Iter() {
				log.Println(imap.Value)
			}
			err = json.Unmarshal([]byte(jsonInput), &newMap)
			if err != nil {
				log.Fatalf("Error unmarshaling JSON: %v", err)
			}
			taskData.Response = ctx.OK(obj)
			taskData.Request = nil
			whenDone(ctx, taskData, nil)
		}))
}

func (r *GetUserByPartnerIdTask) GetName() string {
	return r.Name
}
