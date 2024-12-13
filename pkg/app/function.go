package app

import (
	"fmt"
	"github.com/kimxuanhong/user-manager-go/pkg/utils/ex"
	"log"
	"runtime/debug"
)

type Handler[T any] func(obj T, err error)
type HandlerFunc[T any] func(ctx *Context, whenDone Handler[T])

func panicHandler(whenDone func(err error)) {
	if r := recover(); r != nil {
		log.Printf("Panic recovered: %v\nStack trace: %s", r, string(debug.Stack()))
		whenDone(fmt.Errorf("Recovered from panic: %v\n", r))
		return
	}
}

func SafeCallback[T any](callback Handler[T]) Handler[T] {
	return func(obj T, err error) {
		defer panicHandler(func(er error) {
			callback(obj, ex.New("PANIC_ERROR", er.Error()))
		})
		callback(obj, err)
	}
}

func TryCatch(fnc func(ex error)) {
	defer panicHandler(func(err error) {
		fnc(err)
		return
	})
	fnc(nil)
}
