package app

import (
	"fmt"
	exeption "github.com/kimxuanhong/user-manager-go/pkg/utils/ex"
	"log"
	"runtime/debug"
)

type Handler[T any] func(obj T, err error)
type HandlerFunc[T any] func(ctx *Context, whenDone Handler[T])

type Result[T any] struct {
	Value T
	Error error
}

func panicHandler(whenDone func(err error)) {
	if r := recover(); r != nil {
		log.Printf("Panic recovered: %v\nStack trace: %s", r, string(debug.Stack()))
		whenDone(fmt.Errorf("Recovered from panic: %v\n", r))
		return
	}
}

func SafeCallback[T any](callback Handler[T]) Handler[T] {
	return func(obj T, err error) {
		TryCatch(func(ex error) {
			if ex != nil {
				callback(obj, exeption.New("PANIC_ERROR", ex.Error()))
			}
			callback(obj, err)
		})
	}
}

func TryCatch(fnc func(ex error)) {
	defer panicHandler(func(err error) {
		fnc(err)
		return
	})
	fnc(nil)
}
