package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kimxuanhong/user-manager-go/pkg/api/dto"
	"net/http"
	"reflect"
	"time"
)

type Context struct {
	*gin.Context
	RequestId string
}

type HandlerFunc func(ctx *Context)

var timeLayout = "2006-01-02T15:04:05.000-07:00"

func Handler(handler HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler(&Context{
			Context:   ctx,
			RequestId: uuid.NewString(),
		})
	}
}

func GetRequestId(data interface{}) (interface{}, bool) {
	// Lấy giá trị reflect của struct
	v := reflect.ValueOf(data)
	// Kiểm tra xem data có phải là struct hay không
	if v.Kind() == reflect.Ptr {
		v = v.Elem() // Nếu là pointer, lấy giá trị thực của struct
	}
	if v.Kind() == reflect.Struct {
		// Tìm trường có tên là `fieldName`
		field := v.FieldByName("RequestId")
		// Kiểm tra xem trường đó có tồn tại và có thể lấy giá trị không
		if field.IsValid() {
			return field.Interface(), true
		}
	}
	return nil, false
}

func (ctx *Context) BindAndValidate(data interface{}) error {
	// Bind JSON body to struct
	if err := ctx.ShouldBindJSON(data); err != nil {
		return err
	}
	if id, ok := GetRequestId(data); ok {
		ctx.RequestId = id.(string)
	}
	return nil
}

func (ctx *Context) print(status int, data interface{}) {
	ctx.JSONP(status, data)
	if status != http.StatusOK {
		ctx.Abort()
	}
}

func (ctx *Context) OK(data interface{}) {
	response := &dto.Response{
		ResponseId:   ctx.RequestId,
		ResponseTime: time.Now().Format(timeLayout),
		ResponseCode: dto.SUCCESS,
		Data:         data,
	}
	ctx.print(http.StatusOK, response)
}

func (ctx *Context) Bad(status dto.Status, data interface{}) {
	response := &dto.Response{
		ResponseId:   ctx.RequestId,
		ResponseTime: time.Now().Format(timeLayout),
		ResponseCode: status,
		Data:         data,
	}
	ctx.print(http.StatusBadRequest, response)
}

func (ctx *Context) Error(data interface{}) {
	response := &dto.Response{
		ResponseId:   ctx.RequestId,
		ResponseTime: time.Now().Format(timeLayout),
		ResponseCode: dto.ERROR,
		Data:         data,
	}
	ctx.print(http.StatusInternalServerError, response)
}
