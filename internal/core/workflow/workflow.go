package workflow

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/internal/core/task"
	"github.com/kimxuanhong/user-manager-go/pkg/app"
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
		app.TryCatch(func(ex error) {
			if ex != nil {
				whenDone(ctx, taskData, ex)
				return
			}

			log.Printf("---------------------- Workflow %s starting! ----------------------\n", wf.Name)
			wf.Result = taskData
			for _, taskStep := range wf.Tasks {
				log.Printf("Run %s", taskStep.GetName())
				taskChannel := make(chan struct {
					*task.Data
					error
				}, 1)

				go func(taskStep task.Task) {
					app.TryCatch(func(ex error) {
						if ex != nil {
							taskChannel <- struct {
								*task.Data
								error
							}{taskData, ex}
							return
						}

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
					})
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
		})

	}(ctx, taskData)
}

func NewWorkflow(workflowName string) *Workflow {
	return &Workflow{
		Name:  workflowName,
		Tasks: make([]task.Task, 0),
	}
}
