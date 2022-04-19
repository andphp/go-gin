package result

type InterfaceResult struct {
	Result interface{}
	Err    error
}

func NewInterfaceResult(result interface{}, err error) *InterfaceResult {
	return &InterfaceResult{Result: result, Err: err}
}
func (ain *InterfaceResult) Unwrap() interface{} {
	if ain.Err != nil {
		panic(ain.Err)
	}
	return ain.Result
}
func (ain *InterfaceResult) Unwrap_Or(v interface{}) interface{} {
	if ain.Err != nil {
		return v
	}
	return ain.Result
}
func (ain *InterfaceResult) Unwrap_Or_Else(f func() interface{}) interface{} {
	return f()
}
