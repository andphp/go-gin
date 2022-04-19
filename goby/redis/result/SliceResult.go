package result

type SliceResult struct {
	Result []interface{}
	Err    error
}

func NewSliceResult(result []interface{}, err error) *SliceResult {
	return &SliceResult{Result: result, Err: err}
}
func (ain *SliceResult) Unwrap() []interface{} {
	if ain.Err != nil {
		panic(ain.Err)
	}
	return ain.Result
}
func (ain *SliceResult) Unwrap_Or(v []interface{}) []interface{} {
	if ain.Err != nil {
		return v
	}
	return ain.Result
}
func (ain *SliceResult) Iter() *Iterator {
	return NewIterator(ain.Result)
}
