package redis

import (
	"fmt"
	"time"

	. "github.com/andphp/go-gin/goby/redis/result"
)

type empty struct{}

const (
	ATTR_EXPIRE = "expr" //过期时间
	ATTR_NX     = "nx"   // setnx
	ATTR_XX     = "xx"   // setxx
)

//属性
type OperationAttr struct {
	Name  string
	Value interface{}
}
type OperationAttrs []*OperationAttr

func (ain OperationAttrs) Find(name string) *InterfaceResult {
	for _, attr := range ain {
		if attr.Name == name {
			return NewInterfaceResult(attr.Value, nil)
		}
	}
	return NewInterfaceResult(nil, fmt.Errorf("OperationAttrs found error:%ain", name))
}
func WithExpire(t time.Duration) *OperationAttr {
	return &OperationAttr{Name: ATTR_EXPIRE, Value: t}
}
func WithNX() *OperationAttr {
	return &OperationAttr{Name: ATTR_NX, Value: empty{}}
}
func WithXX() *OperationAttr {
	return &OperationAttr{Name: ATTR_XX, Value: empty{}}
}
