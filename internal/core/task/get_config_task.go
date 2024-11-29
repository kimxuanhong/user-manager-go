package task

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/pkg/api"
	"github.com/kimxuanhong/user-manager-go/pkg/entity"
	"github.com/kimxuanhong/user-manager-go/pkg/libs/task"
	"time"
)

type GetConfigTask struct {
	Name string
}

func NewGetConfigTask() task.Task {
	return &GetConfigTask{
		Name: "GetConfigTask",
	}
}

func (r *GetConfigTask) Execute(ctx *api.Context, taskData *task.Data, whenDone task.Handler) {
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(r.Name + " đang chạy")
		// Giả sử task gặp lỗi
		taskData.Output.(*entity.User).UserName = "Kết quả " + r.Name
		whenDone(ctx, taskData, nil)
	}()
}

func (r *GetConfigTask) GetName() string {
	return r.Name
}
