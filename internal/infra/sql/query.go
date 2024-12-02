package sql

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/pkg/app"
	"github.com/kimxuanhong/user-manager-go/pkg/utils/list"
)

type Params struct {
	Query  string
	Values []interface{}
}

func Query[T any](ctx *app.Context, params Params, whenDone app.Handler[*list.Array[T]]) {
	go func() {
		defer app.PanicHandler(func(obj any, err error) {
			whenDone(list.NewArray[T](), err)
		})
		select {
		case <-ctx.Done():
			whenDone(list.NewArray[T](), fmt.Errorf("context canceled before query execution"))
			return
		default:
			var results []T
			if err := ctx.Db.WithContext(ctx).Raw(params.Query, params.Values...).Scan(&results).Error; err != nil {
				whenDone(list.NewArray[T](), err)
				return
			}
			whenDone(list.AsArray(results), nil)
		}
	}()
}

func Insert[T any](ctx *app.Context, items *list.Array[T], whenDone app.Handler[*list.Array[T]]) {
	go func() {
		defer app.PanicHandler(func(obj any, err error) {
			whenDone(list.NewArray[T](), err)
		})
		select {
		case <-ctx.Done():
			whenDone(list.NewArray[T](), fmt.Errorf("context canceled before query execution"))
			return
		default:
			var results []T
			if err := ctx.Db.WithContext(ctx).Create(items.Slice()).Error; err != nil {
				whenDone(list.NewArray[T](), err)
				return
			}
			whenDone(list.AsArray(results), nil)
		}
	}()
}
