package sql

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/pkg/app"
	"github.com/kimxuanhong/user-manager-go/pkg/utils/list"
	"log"
)

type Params struct {
	Query  string
	Values []interface{}
}

func Query[T any](ctx *app.Context, params Params, whenDone app.Handler[*list.Array[T]]) {
	go func() {
		log.Println("SQL: " + params.Query)
		defer app.PanicHandler(func(obj any, err error) {
			whenDone(list.NewArray[T](), err)
		})
		select {
		case <-ctx.Done():
			whenDone(list.NewArray[T](), fmt.Errorf("context canceled before query execution"))
			return
		default:
			var results []T
			err := ctx.Db.WithContext(ctx).Raw(params.Query, params.Values...).Scan(&results).Error
			if err != nil {
				whenDone(list.NewArray[T](), err)
				return
			}
			whenDone(list.AsArray(results), nil)
		}
	}()
}
