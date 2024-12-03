package app

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/pkg/utils/ex"
	"log"
)

type Handler[T any] func(obj T, err error)
type HandlerFunc[T any] func(ctx *Context, whenDone Handler[T])

func PanicHandler(whenDone func(err error)) {
	if r := recover(); r != nil {
		log.Printf("Recovered from panic: %v\n", r)
		whenDone(fmt.Errorf("Recovered from panic: %v\n", r))
		return
	}
}

func SafeCallback[T any](callback Handler[T]) Handler[T] {
	return func(obj T, err error) {
		defer PanicHandler(func(err error) {
			callback(obj, ex.New("PANIC_ERROR", err.Error()))
		})
		callback(obj, err)
	}
}
