package global

import (
	"go.uber.org/zap"
	"github.com/spf13/viper"
)

var (
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger
)
