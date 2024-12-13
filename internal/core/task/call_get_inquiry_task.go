package task

import (
	"github.com/kimxuanhong/user-manager-go/internal/dto"
	"github.com/kimxuanhong/user-manager-go/internal/infra/entity"
	"github.com/kimxuanhong/user-manager-go/pkg/app"
	"log"
	"time"
)

type CallGetInquiryTask struct {
	Name string
}

func NewCallGetInquiryTask() Task {
	return &CallGetInquiryTask{
		Name: "CallGetInquiryTask",
	}
}

func (r *CallGetInquiryTask) Execute(ctx *app.Context, taskData *Data, whenDone Handler) {
	go func() {
		time.Sleep(1 * time.Second)
		log.Println(r.Name + " đang chạy")
		// Giả sử task gặp lỗi
		taskData.Response.(*dto.Response).Data.(*entity.User).UserName = "Kết quả " + r.Name
		whenDone(ctx, taskData, nil)
	}()
}

func (r *CallGetInquiryTask) GetName() string {
	return r.Name
}
