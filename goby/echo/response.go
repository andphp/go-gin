package echo

import (
	"fmt"

	"github.com/andphp/go-gin/goby/translator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorResult struct {
	data interface{}
	err  error
}

const (
	HTTP_STATE_200 = 200
	HTTP_STATE_401 = 401
	HTTP_STATE_403 = 403
)

func Result(errs ...interface{}) *ErrorResult {
	if len(errs) == 1 {
		if errs[0] == nil {
			return &ErrorResult{nil, nil}
		}
		if e, ok := errs[0].(error); ok {
			return &ErrorResult{nil, e}
		}
	}
	if len(errs) == 2 {
		if errs[1] == nil {
			return &ErrorResult{errs[0], nil}
		}
		if e, ok := errs[1].(error); ok {
			return &ErrorResult{errs[0], e}
		}
	}

	return &ErrorResult{nil, fmt.Errorf("error result format")}
}

func (r *ErrorResult) Unwrap(errType ...ErrorStruct) interface{} {
	if r.err != nil {
		var errError interface{} = r.err.Error()
		// 判断 参数验证
		errs, ok := r.err.(validator.ValidationErrors)
		if ok {
			errError = translator.RemoveTopStruct(errs.Translate(translator.Trans))
		}
		if len(errType) > 0 {
			errType[0].Error = errError
			panic(errType[0])
		} else {
			panic(errError)
		}
	}
	return r.data
}

type ResultFunc func(message string, code int, result interface{}) func(httpState int)

func Output(c *gin.Context) ResultFunc {
	return func(message string, code int, result interface{}) func(httpState int) {
		return func(httpState int) {
			r := ResultPool.Get().(*JSONResult)
			defer ResultPool.Put(r)
			r.Message = message
			r.Code = code
			r.Data = result
			c.AbortWithStatusJSON(httpState, r)
		}
	}
}

func OutputOk(c *gin.Context) {
	Output(c)(SUCCESS.Message, SUCCESS.Code, SUCCESS.Error)(HTTP_STATE_200)
}

func OutputOkMsg(c *gin.Context) func(message string) {
	return func(message string) {
		Output(c)(message, SUCCESS.Code, SUCCESS.Error)(HTTP_STATE_200)
	}
}

func OutputOkData(c *gin.Context) func(data interface{}) {
	return func(data interface{}) {
		Output(c)(SUCCESS.Message, SUCCESS.Code, data)(HTTP_STATE_200)
	}
}

func OutputOkDetail(c *gin.Context) func(data interface{}, message string) {
	return func(data interface{}, message string) {
		Output(c)(message, SUCCESS.Code, data)(HTTP_STATE_200)
	}
}

func OutputFail(c *gin.Context) {
	Output(c)("Failed", ERROR.Code, ERROR.Error)(HTTP_STATE_200)
}

func OutputFailMsg(c *gin.Context) func(message string) {
	return func(message string) {
		Output(c)(message, ERROR.Code, ERROR.Error)(HTTP_STATE_200)
	}
}

func OutputFailDetail(c *gin.Context) func(data interface{}, message string) {
	return func(data interface{}, message string) {
		Output(c)(message, ERROR.Code, data)(HTTP_STATE_200)
	}
}

func OutputFailedSignature(c *gin.Context) func(data interface{}, message string) {
	return func(data interface{}, message string) {
		Output(c)(message, ERROR.Code, data)(HTTP_STATE_401)
	}
}

func OutputFailedPermission(c *gin.Context) func(data interface{}, message string) {
	return func(data interface{}, message string) {
		Output(c)(message, ERROR.Code, data)(HTTP_STATE_403)
	}
}

func OutputError(c *gin.Context) func(error ErrorStruct) {
	return func(error ErrorStruct) {
		Output(c)(error.Message, error.Code, error.Error)(HTTP_STATE_200)
	}
}

func OutputErrorDetail(c *gin.Context) func(data interface{}, error ErrorStruct) {
	return func(data interface{}, error ErrorStruct) {
		Output(c)(error.Message, error.Code, data)(HTTP_STATE_200)
	}
}

func OutputErrorWithHttpState(c *gin.Context) func(error ErrorStruct, httpState int) {
	return func(error ErrorStruct, httpState int) {
		Output(c)(error.Message, error.Code, error.Error)(httpState)
	}
}

func OutputErrorDetailWithHttpState(c *gin.Context) func(data interface{}, error ErrorStruct, httpState int) {
	return func(data interface{}, error ErrorStruct, httpState int) {
		Output(c)(error.Message, error.Code, data)(httpState)
	}
}
