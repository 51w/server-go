package main

import (
	"go.uber.org/zap"
	"fmt"

	"github.com/51w/server-go/server/core"
	"github.com/51w/server-go/server/global"
	"github.com/51w/server-go/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	// initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	} else {
		// global.GVA_LOG.Info("global.GVA_DB is nil!")
		fmt.Println("global.GVA_DB is nil!")
	}
	core.RunWindowsServer()
}
