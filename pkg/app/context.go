package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kimxuanhong/user-manager-go/internal/dto"
	"github.com/kimxuanhong/user-manager-go/pkg/dependencies"
	"net/http"
	"time"
)

type Context struct {
	*gin.Context
	*dependencies.Dependency
	RequestId string
}

func RouteHandler(deps *dependencies.Dependency, handler HandlerFunc[any]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		successChan := make(chan any, 10)
		errorChan := make(chan error, 10)
		defer close(successChan)
		defer close(errorChan)
		go func() {
			TryCatch(func(ex error) {
				if ex != nil {
					errorChan <- fmt.Errorf("internal Server Error. Please try again later")
					return
				}
				handler(&Context{Context: ctx, Dependency: deps, RequestId: uuid.NewString()}, func(obj any, error error) {
					if error != nil {
						errorChan <- error
						return
					}
					successChan <- obj
					return
				})
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

func (ctx *Context) OK(data any) *dto.Response {
	return &dto.Response{
		ResponseId:   ctx.RequestId,
		ResponseTime: time.Now().Format(timeLayout),
		ResponseCode: dto.SUCCESS,
		Data:         data,
	}
}

func (ctx *Context) Bad(status dto.Status, data any) *dto.Response {
	return &dto.Response{
		ResponseId:   ctx.RequestId,
		ResponseTime: time.Now().Format(timeLayout),
		ResponseCode: status,
		Data:         data,
	}
}

func (ctx *Context) Error(data any) *dto.Response {
	return &dto.Response{
		ResponseId:   ctx.RequestId,
		ResponseTime: time.Now().Format(timeLayout),
		ResponseCode: dto.ERROR,
		Data:         data,
	}
}
