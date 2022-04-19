package redis

import (
	"context"
	"time"

	"github.com/andphp/go-gin/goby"
	. "github.com/andphp/go-gin/goby/redis/result"
)

//专门处理string类型的操作
type StringOperation struct {
	ctx context.Context
}

func NewStringOperation() *StringOperation {
	return &StringOperation{ctx: context.Background()}
}

func (ain *StringOperation) Set(key string, value interface{},
	attrs ...*OperationAttr) *InterfaceResult {
	exp := OperationAttrs(attrs).
		Find(ATTR_EXPIRE).
		Unwrap_Or(time.Second * 0).(time.Duration)

	nx := OperationAttrs(attrs).Find(ATTR_NX).Unwrap_Or(nil)
	if nx != nil {
		return NewInterfaceResult(goby.GOBY_REDIS.SetNX(ain.ctx, key, value, exp).Result())
	}
	xx := OperationAttrs(attrs).Find(ATTR_XX).Unwrap_Or(nil)
	if xx != nil {
		return NewInterfaceResult(goby.GOBY_REDIS.SetXX(ain.ctx, key, value, exp).Result())
	}
	return NewInterfaceResult(goby.GOBY_REDIS.Set(ain.ctx, key, value, exp).Result())

}
func (ain *StringOperation) Get(key string) *StringResult {
	return NewStringResult(goby.GOBY_REDIS.Get(ain.ctx, key).Result())
}
func (ain *StringOperation) MGet(keys ...string) *SliceResult {
	return NewSliceResult(goby.GOBY_REDIS.MGet(ain.ctx, keys...).Result())
}
