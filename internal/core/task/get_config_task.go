package task

import (
	"github.com/kimxuanhong/user-manager-go/internal/dto"
	"github.com/kimxuanhong/user-manager-go/internal/infra/entity"
	"github.com/kimxuanhong/user-manager-go/pkg/app"
	"log"
	"time"
)

type GetConfigTask struct {
	Name string
}

func NewGetConfigTask() Task {
	return &GetConfigTask{
		Name: "GetConfigTask",
	}
}

func (r *GetConfigTask) Execute(ctx *app.Context, taskData *Data, whenDone Handler) {
	go func() {
		app.TryCatch(func(ex error) {
			if ex != nil {
				whenDone(ctx, taskData, ex)
				return
			}
			time.Sleep(1 * time.Second)
			log.Println(r.Name + " đang chạy")
			// Giả sử task gặp lỗi
			taskData.Response.(*dto.Response).Data.(*entity.User).UserName = "Kết quả " + r.Name
			whenDone(ctx, taskData, nil)
		})
	}()
}

func (r *GetConfigTask) GetName() string {
	return r.Name
}
