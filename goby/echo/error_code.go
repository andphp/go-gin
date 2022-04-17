package echo

type ErrorStruct struct {
	Code    int
	Message string
	Error   interface{}
}

var (
	SUCCESS                = ErrorStruct{0, "Successful", []interface{}{}}
	ERROR                  = ErrorStruct{7, "Failed", []interface{}{}}
	SYSTEM_ERROR           = ErrorStruct{100000, "系统错误", []interface{}{}}
	PARAMS_ERROR           = ErrorStruct{100001, "参数错误", []interface{}{}}
	ACCOUNT_PASSWORD_ERROR = ErrorStruct{100020, "账号密码错误", []interface{}{}}
)
