package result

type Iterator struct {
	data  []interface{}
	index int
}

func NewIterator(data []interface{}) *Iterator {
	return &Iterator{data: data}
}
func (ain *Iterator) HasNext() bool {
	if ain.data == nil || len(ain.data) == 0 {
		return false
	}
	return ain.index < len(ain.data)
}
func (ain *Iterator) Next() (ret interface{}) {
	ret = ain.data[ain.index]
	ain.index = ain.index + 1
	return
}
