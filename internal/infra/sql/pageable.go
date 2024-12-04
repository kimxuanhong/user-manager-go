package sql

import (
	"github.com/kimxuanhong/user-manager-go/pkg/app"
	"github.com/kimxuanhong/user-manager-go/pkg/utils/list"
	"strings"
)

type PageData interface {
	GetTotal() int
}

type Pageable[T PageData] struct {
	query       string
	values      []interface{}
	whereClause strings.Builder
	pageNumber  int
	pageSize    int
}

type Page[T PageData] struct {
	Data         *list.Array[T]
	PageNumber   int
	PageSize     int
	TotalElement int
	TotalPage    int
}

func InitPage[T PageData]() *Pageable[T] {
	return &Pageable[T]{
		values:     make([]interface{}, 0),
		pageNumber: 1,
		pageSize:   30,
	}
}

func (f *Pageable[T]) AndWhere(clause string, param ...interface{}) *Pageable[T] {
	if len(param) > 0 {
		for _, p := range param {
			if p == nil {
				return f
			}
		}
		f.whereClause.WriteString(" AND " + clause)
		f.values = append(f.values, param...)
	}
	return f
}

func (f *Pageable[T]) OrWhere(clause string, param ...interface{}) *Pageable[T] {
	if len(param) > 0 {
		for _, p := range param {
			if p == nil {
				return f
			}
		}
		f.whereClause.WriteString(" OR " + clause)
		f.values = append(f.values, param...)
	}
	return f
}

func (f *Pageable[T]) SetQuery(query string) *Pageable[T] {
	f.query = query
	return f
}

func (f *Pageable[T]) SetPageNumber(pageNumber int) *Pageable[T] {
	if (pageNumber - 1) <= 0 {
		return f
	}
	f.pageNumber = pageNumber
	return f
}

func (f *Pageable[T]) SetPageSize(pageSize int) *Pageable[T] {
	if pageSize <= 0 {
		return f
	}
	f.pageSize = pageSize
	return f
}

func (f *Pageable[T]) GetSql() string {
	var builder strings.Builder
	builder.WriteString("\nWITH temp_tbl AS (")
	builder.WriteString("\n" + f.query + " " + f.whereClause.String())
	builder.WriteString("\n),")
	builder.WriteString("\ncte_total AS (SELECT COUNT(*) AS all_element FROM temp_tbl)")
	builder.WriteString("\nSELECT  tmp.*, cte.all_element FROM cte_total AS cte, temp_tbl AS tmp WHERE 1=1")
	builder.WriteString("\nLIMIT ? OFFSET ?")
	return builder.String()
}

func (f *Pageable[T]) GetParams() []interface{} {
	return f.values
}

func (f *Pageable[T]) GetLimit() int {
	return f.pageSize
}

func (f *Pageable[T]) GetOffset() int {
	return f.pageSize * (f.pageNumber - 1)
}

func (f *Pageable[T]) Fetch(ctx *app.Context, whenDone app.Handler[*Page[T]]) {
	params := append(f.GetParams(), f.GetLimit(), f.GetOffset())
	Query(ctx, Params{Query: f.GetSql(), Values: params}, app.SafeCallback(func(obj *list.Array[T], err error) {
		if err != nil {
			whenDone(&Page[T]{
				PageNumber: f.pageNumber,
				PageSize:   f.pageSize,
			}, err)
			return
		}
		total := getTotal(obj)
		whenDone(&Page[T]{
			Data:         obj,
			PageNumber:   f.pageNumber,
			PageSize:     f.pageSize,
			TotalElement: total,
			TotalPage:    total / f.pageSize,
		}, err)
	}))
}

func getTotal[T PageData](obj *list.Array[T]) int {
	if first, err := obj.First(); err != nil {
		return 0
	} else {
		return first.GetTotal()
	}
}
