package task

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/pkg/api"
	"github.com/kimxuanhong/user-manager-go/pkg/entity"
	"github.com/kimxuanhong/user-manager-go/pkg/libs/task"
	"time"
)

type CallGetInquiryTask struct {
	Name string
}

func NewCallGetInquiryTask() task.Task {
	return &CallGetInquiryTask{
		Name: "CallGetInquiryTask",
	}
}

func (r *CallGetInquiryTask) Execute(ctx *api.Context, taskData *task.Data, whenDone task.Handler) {
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(r.Name + " đang chạy")
		// Giả sử task gặp lỗi
		taskData.Output.(*entity.User).UserName = "Kết quả " + r.Name
		whenDone(ctx, taskData, nil)
	}()
}
