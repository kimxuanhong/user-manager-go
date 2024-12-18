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

func Query[T any](ctx *app.Context, params Params, whenDone app.Handler[*list.List[T]]) {
	go func() {
		app.TryCatch(func(ex error) {
			if ex != nil {
				whenDone(list.NewList[T](), ex)
				return
			}

			select {
			case <-ctx.Done():
				whenDone(list.NewList[T](), fmt.Errorf("context canceled before query execution"))
				return
			default:
				var results []T
				if err := ctx.Db.WithContext(ctx).Raw(params.Query, params.Values...).Scan(&results).Error; err != nil {
					whenDone(list.NewList[T](), err)
					return
				}
				whenDone(list.AsList(results...), nil)
			}
		})
	}()
}

func Insert[T any](ctx *app.Context, items *list.List[T], whenDone app.Handler[*list.List[T]]) {
	go func() {
		app.TryCatch(func(ex error) {
			if ex != nil {
				whenDone(list.NewList[T](), ex)
				return
			}

			select {
			case <-ctx.Done():
				whenDone(list.NewList[T](), fmt.Errorf("context canceled before query execution"))
				return
			default:
				var results []T
				if err := ctx.Db.WithContext(ctx).Create(items.Slice()).Error; err != nil {
					whenDone(list.NewList[T](), err)
					return
				}
				whenDone(list.AsList(results...), nil)
			}
		})
	}()
}
