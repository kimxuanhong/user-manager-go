package task

import (
	"github.com/kimxuanhong/user-manager-go/internal/infra/entity"
	"github.com/kimxuanhong/user-manager-go/pkg/app"
	"github.com/kimxuanhong/user-manager-go/pkg/task"
	"log"
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

func (r *CacheConfigTask) Execute(ctx *app.Context, taskData *task.Data, whenDone task.Handler) {
	go func() {
		time.Sleep(1 * time.Second)
		log.Println(r.Name + " đang chạy")
		// Giả sử task gặp lỗi
		taskData.Output.(*entity.User).UserName = "Kết quả " + r.Name
		whenDone(ctx, taskData, nil)
	}()
}

func (r *CacheConfigTask) GetName() string {
	return r.Name
}
