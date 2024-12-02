package app

import (
	"fmt"
	"log"
)

type Handler[T any] func(obj T, err error)
type HandlerFunc[T any] func(ctx *Context, whenDone Handler[T])

func PanicHandler(whenDone Handler[any]) {
	if r := recover(); r != nil {
		log.Printf("Recovered from panic: %v\n", r)
		whenDone(nil, fmt.Errorf("Recovered from panic: %v\n", r))
		return
	}
}
