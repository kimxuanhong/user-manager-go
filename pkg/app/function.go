package app

import (
	"fmt"
	"log"
)

type Handler[T any] func(obj T, error error)
type HandlerFunc[T any] func(ctx *Context, whenDone Handler[T])

func SafeGo(whenDone Handler[any], fn func()) {
	func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v\n", r)
				whenDone(nil, fmt.Errorf("Recovered from panic: %v\n", r))
				return
			}
		}()
		fn()
	}()
}
