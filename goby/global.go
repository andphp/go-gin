package goby

import (
	conf "github.com/andphp/go-gin/goby/config"
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"

	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	GOBY_CONFIG conf.Server
	GOBY_VIPER  *viper.Viper
	GOBY_LOG    *zap.Logger
	GOBY_DB     *gorm.DB
	GOBY_REDIS  *redis.Client
	GOBY_BlackCache  local_cache.Cache
	GOBY_Concurrency_Control = &singleflight.Group{}
)
