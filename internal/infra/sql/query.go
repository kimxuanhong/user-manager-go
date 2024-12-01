package sql

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/pkg/app"
)

type Params struct {
	Query  string
	Values []interface{}
}

func QueryWithParams[T any](ctx *app.Context, db *Database, params Params, whenDone app.Handler[[]T]) {
	go func() {
		select {
		case <-ctx.Done():
			whenDone(nil, fmt.Errorf("context canceled before query execution"))
			return
		default:
			var results []T
			err := db.DB.WithContext(ctx).Raw(params.Query, params.Values...).Scan(&results).Error
			if err != nil {
				whenDone(nil, err)
				return
			}
			whenDone(results, nil)
		}
	}()
}

func QueryWithoutParams[T any](ctx *app.Context, db *Database, query string, whenDone app.Handler[[]T]) {
	go func() {
		select {
		case <-ctx.Done():
			whenDone(nil, fmt.Errorf("context canceled before query execution"))
			return
		default:
			var results []T
			err := db.DB.WithContext(ctx).Raw(query).Scan(&results).Error
			if err != nil {
				whenDone(nil, err)
				return
			}
			whenDone(results, nil)
		}
	}()
}
