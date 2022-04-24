package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/andphp/go-gin/goby"
	"github.com/go-redis/redis/v8"
)

type Locker struct {
	key        string
	expire     time.Duration
	unlock     bool
	incrScript *redis.Script
}

const incrLua = `
if redis.call('get', KEYS[1]) == ARGV[1] then
  return redis.call('expire', KEYS[1],ARGV[2]) 				
 else
   return '0' 					
end`

func NewLocker(key string) *Locker {
	//默认30秒过期时间
	return &Locker{key: key, expire: time.Second * 30, incrScript: redis.NewScript(incrLua)} //默认30秒
}

//有过期时间
func NewLockerWithTTL(key string, expire time.Duration) *Locker {
	if expire.Seconds() <= 0 {
		panic("error expire")
	}
	return &Locker{key: key, expire: expire, incrScript: redis.NewScript(incrLua)}
}
func (l *Locker) Lock() *Locker {
	boolcmd := goby.GOBY_REDIS.SetNX(context.Background(), l.key, "1", l.expire)
	if ok, err := boolcmd.Result(); err != nil || !ok {
		panic(fmt.Sprintf("lock error with key:%s", l.key))
	}
	l.expandLockTime()
	return l
}
func (l *Locker) expandLockTime() {
	sleepTime := l.expire.Seconds() * 2 / 3
	go func() {
		for {
			time.Sleep(time.Second * time.Duration(sleepTime))
			if l.unlock {
				break
			}
			l.resetExpire()
		}
	}()
}

//重新设置过期时间
func (l *Locker) resetExpire() {
	cmd := l.incrScript.Run(context.Background(), goby.GOBY_REDIS, []string{l.key}, 1, l.expire.Seconds())
	v, err := cmd.Result()
	log.Printf("key=%s ,续期结果:%v,%v\n", l.key, err, v)
}

func (l *Locker) Unlock() {
	l.unlock = true
	goby.GOBY_REDIS.Del(context.Background(), l.key)
}
