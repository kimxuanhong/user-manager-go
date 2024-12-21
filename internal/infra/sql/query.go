package sql

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/pkg/app"
	"github.com/kimxuanhong/user-manager-go/pkg/async"
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

func Select[T any](ctx *app.Context, params Params) <-chan app.Result[*list.List[T]] {
	result := make(chan app.Result[*list.List[T]])
	go func() {
		defer close(result)
		app.TryCatch(func(ex error) {
			if ex != nil {
				result <- app.Result[*list.List[T]]{
					Value: list.NewList[T](),
					Error: ex,
				}
				return
			}

			select {
			case <-ctx.Done():
				result <- app.Result[*list.List[T]]{
					Value: list.NewList[T](),
					Error: fmt.Errorf("context canceled before query execution"),
				}
				return
			default:
				var results []T
				if err := ctx.Db.WithContext(ctx).Raw(params.Query, params.Values...).Scan(&results).Error; err != nil {
					result <- app.Result[*list.List[T]]{
						Value: list.NewList[T](),
						Error: err,
					}

					return
				}
				result <- app.Result[*list.List[T]]{
					Value: list.AsList(results...),
					Error: nil,
				}
			}
		})
	}()

	return result
}

func Query2[T any](ctx *app.Context, params Params) async.Future[*list.List[T]] {
	return async.Promise(func() (*list.List[T], error) {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context canceled before query execution")
		default:
			var results []T
			if err := ctx.Db.WithContext(ctx).Raw(params.Query, params.Values...).Scan(&results).Error; err != nil {
				return list.NewList[T](), err
			}
			return list.AsList(results...), nil
		}
	})
}

func Query3[T any](ctx *app.Context, params Params) async.Future[*list.List[T]] {
	return async.Promise(func() (*list.List[T], error) {
		promise := Query2[T](ctx, params)
		return promise.Await()
	})
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
