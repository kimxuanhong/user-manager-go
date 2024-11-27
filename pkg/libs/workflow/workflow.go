package workflow

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/pkg/api"
	"github.com/kimxuanhong/user-manager-go/pkg/libs/task"
	"log"
)

type Workflow struct {
	Tasks  []task.Task
	Result *task.Data
	Error  error
}

func (wf *Workflow) AddTask(task task.Task) {
	wf.Tasks = append(wf.Tasks, task)
}

func (wf *Workflow) Run(ctx *api.Context, taskData *task.Data, whenDone task.Handler) {
	go func(ctx *api.Context, taskData *task.Data) {
		log.Println("---------------------- Workflow starting! ----------------------")
		wf.Result = taskData
		for _, taskStep := range wf.Tasks {
			taskChannel := make(chan struct {
				*task.Data
				error
			}, 1)

			go func(taskStep task.Task) {
				select {
				case <-ctx.Done():
					taskChannel <- struct {
						*task.Data
						error
					}{wf.Result, fmt.Errorf("context cancled")}
				default:
					taskStep.Execute(ctx, wf.Result, func(ctx *api.Context, result *task.Data, err error) {
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
				log.Printf("---------------------- Task failed: %v\n", result.error)
				wf.Error = result.error
				wf.Result = result.Data
				whenDone(ctx, wf.Result, wf.Error)
				return
			}
			log.Println("---------------------- Task success ----------------------")
			wf.Result = result.Data
		}

		log.Println("---------------------- Workflow success! ----------------------")
		whenDone(ctx, wf.Result, wf.Error)
	}(ctx, taskData)
}

func NewWorkflow() *Workflow {
	return &Workflow{
		Tasks: make([]task.Task, 0),
	}
}
