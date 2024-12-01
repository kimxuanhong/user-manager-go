package app

type Handler[T any] func(obj T, error error)
type HandlerFunc[T any] func(ctx *Context, whenDone Handler[T])
