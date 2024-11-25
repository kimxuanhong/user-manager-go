package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kimxuanhong/user-manager-go/pkg/api/config"
	"net/http"
	"time"
)

type Context struct {
	*gin.Context
	Deps      *config.Dependencies
	RequestId string
}

type Handler[T any] func(obj T, error error)
type HandlerFunc[T any] func(ctx *Context, whenDone Handler[T])

func RouteHandler(deps *config.Dependencies, handler HandlerFunc[any]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		successChan := make(chan any, 1)
		errorChan := make(chan error, 1)
		defer close(successChan)
		defer close(errorChan)
		go func() {
			handler(&Context{Context: ctx, Deps: deps, RequestId: uuid.NewString()}, func(obj any, error error) {
				if error != nil {
					errorChan <- error
					return
				}
				successChan <- obj
				return
			})
		}()

		select {
		case res := <-successChan:
			ctx.JSON(http.StatusOK, res)
			return
		case err := <-errorChan:
			ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		case <-ctx.Done():
			ctx.JSON(http.StatusOK, gin.H{"error": "request canceled"})
			return
		}
	}
}

var timeLayout = "2006-01-02T15:04:05.000-07:00"

func (ctx *Context) Bind(obj any) error {
	if err := ctx.ShouldBindJSON(&obj); err != nil {
		return err
	}
	return nil
}

func (ctx *Context) SetRequestId(requestId string) {
	ctx.RequestId = requestId
}

func (ctx *Context) OK(data any) *Response {
	return &Response{
		ResponseId:   ctx.RequestId,
		ResponseTime: time.Now().Format(timeLayout),
		ResponseCode: SUCCESS,
		Data:         data,
	}
}

func (ctx *Context) Bad(status Status, data any) *Response {
	return &Response{
		ResponseId:   ctx.RequestId,
		ResponseTime: time.Now().Format(timeLayout),
		ResponseCode: status,
		Data:         data,
	}
}

func (ctx *Context) Error(data any) *Response {
	return &Response{
		ResponseId:   ctx.RequestId,
		ResponseTime: time.Now().Format(timeLayout),
		ResponseCode: ERROR,
		Data:         data,
	}
}
