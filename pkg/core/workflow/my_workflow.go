package workflow

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/pkg/api/api"
	"github.com/kimxuanhong/user-manager-go/pkg/infra/entity"
	"time"
)

type MyWorkFlow struct {
	*Workflow
}

func NewMyWorkFlow() *MyWorkFlow {
	wf := &MyWorkFlow{
		Workflow: NewWorkflow(),
	}
	wf.AddTask("Task1", wf.Task1)
	wf.AddTask("Task2", wf.Task2)
	wf.AddTask("Task3", wf.Task3)
	wf.AddTask("Task4", wf.Task4)
	return wf
}

func (r *MyWorkFlow) Task4(ctx *api.Context, taskData *TaskData, callback Handler) {
	go func() {
		fmt.Println("Task4 đang chạy")
		// Giả sử task thành công
		taskData.Output.(*entity.User).UserName = "Kết quả Task4"
		callback(ctx, taskData, nil)
	}()

}

func (r *MyWorkFlow) Task3(ctx *api.Context, taskData *TaskData, callback Handler) {

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Task3 đang chạy")
		// Giả sử task thành công
		taskData.Output.(*entity.User).UserName = "Kết quả Task3"
		callback(ctx, taskData, fmt.Errorf("loi ne"))
	}()

}

func (r *MyWorkFlow) Task2(ctx *api.Context, taskData *TaskData, callback Handler) {

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Task2 đang chạy")
		// Giả sử task gặp lỗi
		taskData.Output.(*entity.User).UserName = "Kết quả Task2"
		callback(ctx, taskData, nil)
	}()

}

func (r *MyWorkFlow) Task1(ctx *api.Context, taskData *TaskData, callback Handler) {

	time.Sleep(2 * time.Second)
	fmt.Println("Task1 đang chạy")
	// Giả sử task thành công
	taskData.Output.(*entity.User).UserName = "Kết quả Task1"
	taskData.Input.(*entity.User).UserName = "Kết quả Task1"
	callback(ctx, taskData, nil)

}
