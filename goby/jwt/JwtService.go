package jwt

import (
	"context"
	"github.com/andphp/go-gin/goby"
	"github.com/andphp/go-gin/goby/redis"
	"time"
)

type JwtService struct {
}

//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: token string
//@return: err error

func (jwtService *JwtService) JsonInBlacklist(token string) (err error) {
	goby.BlackCache.SetDefault(token, struct{}{})
	return
}

//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	ok := redis.NewStringOperation().Get(jwt).Unwrap_Or("")
	if ok != "" {
		return true
	}
	return false
}

//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: err error, redisJWT string

func (jwtService *JwtService) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = goby.GOBY_REDIS.Get(context.Background(), userName).Result()
	return err, redisJWT
}

//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(goby.GOBY_CONFIG.JWT.ExpiresTime) * time.Second
	err = goby.GOBY_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}
