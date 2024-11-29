package task

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/pkg/api"
	"github.com/kimxuanhong/user-manager-go/pkg/entity"
	"github.com/kimxuanhong/user-manager-go/pkg/libs/task"
	"time"
)

type CacheConfigTask struct {
	Name string
}

func NewCacheConfigTask() task.Task {
	return &CacheConfigTask{
		Name: "CacheConfigTask",
	}
}

func (r *CacheConfigTask) Execute(ctx *api.Context, taskData *task.Data, whenDone task.Handler) {
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(r.Name + " đang chạy")
		// Giả sử task gặp lỗi
		taskData.Output.(*entity.User).UserName = "Kết quả " + r.Name
		whenDone(ctx, taskData, nil)
	}()
}

func (r *CacheConfigTask) GetName() string {
	return r.Name
}
