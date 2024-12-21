package async

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"
	"sync"
)

type Future[T any] interface {
	Await() (T, error)
	AwaitWithCallback(callback func(result T, err error))
}

type future[T any] struct {
	await func(ctx context.Context) (T, error)
}

func (f future[T]) Await() (T, error) {
	return f.await(context.Background())
}

func (f future[T]) AwaitWithCallback(callback func(result T, err error)) {
	result, err := f.await(context.Background())
	callback(result, err)
}

func Promise[T any](f func() (T, error)) Future[T] {
	resultChan := make(chan struct {
		result T
		err    error
	}, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic recovered: %v\nStack trace: %s", r, string(debug.Stack()))
				var defaultResult T
				resultChan <- struct {
					result T
					err    error
				}{defaultResult, fmt.Errorf("panic recovered: %v", r)}
				return
			}
		}()

		defer close(resultChan)
		result, err := f()
		resultChan <- struct {
			result T
			err    error
		}{result, err}

	}()

	return future[T]{
		await: func(ctx context.Context) (T, error) {
			var defaultResult T
			select {
			case <-ctx.Done():
				return defaultResult, ctx.Err()
			case res := <-resultChan:
				return res.result, res.err
			}
		},
	}
}

func OfAll[T any](futures ...Future[T]) ([]T, []error) {
	var wg sync.WaitGroup
	results := make([]T, len(futures))
	errs := make([]error, len(futures))
	for i, f := range futures {
		wg.Add(1)
		go func(i int, f Future[T]) {
			defer wg.Done()
			result, err := f.Await()
			results[i] = result
			errs[i] = err
		}(i, f)
	}
	wg.Wait()
	return results, errs
}
