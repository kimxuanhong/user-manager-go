package workflow

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/pkg/app"
	"github.com/kimxuanhong/user-manager-go/pkg/task"
	"log"
)

type Workflow struct {
	Name   string
	Tasks  []task.Task
	Result *task.Data
	Error  error
}

func (wf *Workflow) AddTask(task task.Task) {
	wf.Tasks = append(wf.Tasks, task)
}

func (wf *Workflow) Run(ctx *app.Context, taskData *task.Data, whenDone task.Handler) {
	go func(ctx *app.Context, taskData *task.Data) {
		defer app.PanicHandler(func(obj any, err error) {
			whenDone(ctx, taskData, err)
		})
		log.Printf("---------------------- Workflow %s starting! ----------------------\n", wf.Name)
		wf.Result = taskData
		for _, taskStep := range wf.Tasks {
			log.Printf("Run %s", taskStep.GetName())
			taskChannel := make(chan struct {
				*task.Data
				error
			}, 1)

			go func(taskStep task.Task) {
				defer app.PanicHandler(func(obj any, err error) {
					taskChannel <- struct {
						*task.Data
						error
					}{nil, err}
				})
				select {
				case <-ctx.Done():
					taskChannel <- struct {
						*task.Data
						error
					}{wf.Result, fmt.Errorf("context cancled")}
				default:
					taskStep.Execute(ctx, wf.Result, func(ctx *app.Context, result *task.Data, err error) {
						taskChannel <- struct {
							*task.Data
							error
						}{result, err}
					})
				}
			}(taskStep)

			result := <-taskChannel
			close(taskChannel)

			if result.error != nil {
				wf.Error = result.error
				wf.Result = result.Data
				whenDone(ctx, wf.Result, wf.Error)
				return
			}
			wf.Result = result.Data
		}

		log.Printf("---------------------- Workflow %s success! ----------------------\n", wf.Name)
		whenDone(ctx, wf.Result, wf.Error)
	}(ctx, taskData)
}

func NewWorkflow(workflowName string) *Workflow {
	return &Workflow{
		Name:  workflowName,
		Tasks: make([]task.Task, 0),
	}
}
