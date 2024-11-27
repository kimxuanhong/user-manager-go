package workflow

import (
	"github.com/kimxuanhong/user-manager-go/internal/core/task"
	"github.com/kimxuanhong/user-manager-go/pkg/libs/workflow"
)

type MyWorkFlow struct {
	*workflow.Workflow
}

func NewMyWorkFlow() *MyWorkFlow {
	wf := &MyWorkFlow{
		Workflow: workflow.NewWorkflow(),
	}
	wf.AddTask(task.NewGetConfigTask())
	wf.AddTask(task.NewCacheConfigTask())
	wf.AddTask(task.NewCallGetInquiryTask())
	return wf
}
