package dao

import (
	"github.com/kimxuanhong/user-manager-go/pkg/api/api"
	"github.com/kimxuanhong/user-manager-go/pkg/core/workflow"
	"github.com/kimxuanhong/user-manager-go/pkg/infra/entity"
	"github.com/kimxuanhong/user-manager-go/pkg/infra/sql"
	"log"
	"sync"
)

type UserDao interface {
	FindUserByPartnerId(ctx *api.Context, partnerId string, whenDone api.Handler[[]entity.User])
}

type userDao struct {
}

var instanceUserDao *userDao
var userDaoOnce sync.Once

func NewUserDao() UserDao {
	userDaoOnce.Do(func() {
		instanceUserDao = &userDao{}
	})
	return instanceUserDao
}

func (r *userDao) FindUserByPartnerId(ctx *api.Context, partnerId string, whenDone api.Handler[[]entity.User]) {

	wf := workflow.NewMyWorkFlow()
	// Khởi tạo taskData ban đầu
	taskData := &workflow.TaskData{
		Input:  &entity.User{},
		Output: &entity.User{},
	}
	wf.Run(ctx, taskData, func(ctx *api.Context, taskData *workflow.TaskData, err error) {
		log.Println("OUTPUT: " + taskData.Output.(*entity.User).UserName)
		QueryWithParams(ctx, Params{Query: sql.GetUserByPartnerId, Values: []interface{}{partnerId}}, func(users []entity.User, err error) {
			if err != nil {
				log.Println("Query was error!")
				whenDone(nil, err)
				return
			}
			whenDone(users, nil)
		})
	})
}
