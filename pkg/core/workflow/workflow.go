package workflow

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/pkg/api/api"
)

type TaskData struct {
	Input  interface{}
	Output interface{}
}

type Handler func(ctx *api.Context, taskData *TaskData, err error)

type Task struct {
	Name    string
	Execute func(ctx *api.Context, taskData *TaskData, whenDone Handler)
}

type Workflow struct {
	Tasks  []*Task
	Result *TaskData
	Error  error
}

func (wf *Workflow) AddTask(name string, execute func(ctx *api.Context, taskData *TaskData, whenDone Handler)) {
	wf.Tasks = append(wf.Tasks, &Task{
		Name:    name,
		Execute: execute,
	})
}

func (wf *Workflow) Run(ctx *api.Context, taskData *TaskData, whenDone Handler) {
	go func(ctx *api.Context, taskData *TaskData) {
		wf.Result = taskData
		for _, task := range wf.Tasks {
			taskChannel := make(chan struct {
				*TaskData
				error
			}, 1)

			go func(task *Task) {
				select {
				case <-ctx.Done():
					taskChannel <- struct {
						*TaskData
						error
					}{wf.Result, fmt.Errorf("context cancled")}
				default:
					task.Execute(ctx, wf.Result, func(ctx *api.Context, result *TaskData, err error) {
						taskChannel <- struct {
							*TaskData
							error
						}{result, err}
					})
				}

			}(task)

			result := <-taskChannel
			close(taskChannel)

			if result.error != nil {
				fmt.Printf("---------------------- Task %s failed ----------------------\n", task.Name)
				fmt.Printf("---------------------- Task %s failed reason: %v\n", task.Name, result.error)
				wf.Error = result.error
				wf.Result = result.TaskData
				whenDone(ctx, wf.Result, wf.Error)
				return
			}
			fmt.Printf("---------------------- Task %s success ----------------------\n", task.Name)
			wf.Result = result.TaskData
		}

		fmt.Println("---------------------- Workflow success! ----------------------")
		whenDone(ctx, wf.Result, wf.Error)
	}(ctx, taskData)
}

func NewWorkflow() *Workflow {
	return &Workflow{
		Tasks: make([]*Task, 0),
	}
}
