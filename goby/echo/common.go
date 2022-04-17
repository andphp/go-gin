package echo

import "sync"

var ResultPool *sync.Pool

func init() {
	ResultPool = &sync.Pool{
		New: func() interface{} {
			return NewJSONResult("", 0, nil)
		},
	}
}

func NewJSONResult(message string, code int, result interface{}) *JSONResult {
	return &JSONResult{Message: message, Code: code, Data: result}
}

type JSONResult struct {
	Message string      `json:"msg"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}
type PageResult struct {
	List     interface{} `json:"data"`
	Total    int64       `json:"total"`
	Page     uint64      `json:"page"`
	PageSize uint64      `json:"pageSize"`
	Success  bool        `json:"success"`
}
