package workflow

import (
	"github.com/kimxuanhong/user-manager-go/internal/core/task"
)

type MyWorkFlow struct {
	*Workflow
}

func NewMyWorkFlow() *MyWorkFlow {
	wf := &MyWorkFlow{
		Workflow: NewWorkflow("MyWorkFlow"),
	}
	wf.AddTask(task.NewGetUserByPartnerIdTask())
	//wf.AddTask(task.NewGetConfigTask())
	//wf.AddTask(task.NewCacheConfigTask())
	//wf.AddTask(task.NewCallGetInquiryTask())
	return wf
}
