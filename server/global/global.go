package global

import (
	"sync"

	"go.uber.org/zap"

	"github.com/51w/server-go/server/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_DBList map[string] *gorm.DB

	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger

	lock       sync.RWMutex
)
