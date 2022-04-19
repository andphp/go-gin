package result

type Result interface {
	Unwrap()
}
type StringResult struct {
	Result string
	Err    error
}

func NewStringResult(result string, err error) *StringResult {
	return &StringResult{Result: result, Err: err}
}
func (ain *StringResult) Unwrap() string {
	if ain.Err != nil {
		panic(ain.Err)
	}
	return ain.Result
}
func (ain *StringResult) Unwrap_Or(str string) string {
	if ain.Err != nil {
		return str
	}
	return ain.Result
}
func (ain *StringResult) Unwrap_Or_Else(f func() string) string {
	if ain.Err != nil {
		return f()
	}
	return ain.Result
}
