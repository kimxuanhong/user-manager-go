package task

import (
	"github.com/kimxuanhong/user-manager-go/pkg/api"
)

type Data struct {
	Input  interface{}
	Output interface{}
}

type Handler func(ctx *api.Context, taskData *Data, err error)

type Task interface {
	Execute(ctx *api.Context, taskData *Data, whenDone Handler)
}
