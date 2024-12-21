package task

import (
	"github.com/kimxuanhong/user-manager-go/pkg/app"
)

type Data struct {
	Request  interface{}
	Response interface{}
}

type Handler func(ctx *app.Context, taskData *Data, err error)

type Task interface {
	Execute(ctx *app.Context, taskData *Data, whenDone Handler)
	GetName() string
}
