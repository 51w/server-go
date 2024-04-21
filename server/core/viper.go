package core

import (
	"github.com/spf13/viper"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
// Author [SliverHorn](https://github.com/SliverHorn)
func Viper(path ...string) *viper.Viper {
	

	v := viper.New()
	// v.SetConfigFile(config)
	v.SetConfigType("yaml")

	return v
}
