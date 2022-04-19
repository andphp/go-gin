package goby

import (
	conf "github.com/andphp/go-gin/goby/config"
	"github.com/spf13/viper"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GOBY_CONFIG conf.Server
	GOBY_VIPER  *viper.Viper
	GOBY_LOG    *zap.Logger
	GOBY_DB     *gorm.DB
)
