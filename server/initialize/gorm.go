package initialize

import (
	"os"

	"github.com/51w/server-go/server/global"
	"github.com/51w/server-go/server/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "sqlite":
		return GormSqlite()
	default:
		return GormMysql()
	}
}

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Name     string `gorm:"size:255"`
}

func RegisterTables() {
	db := global.GVA_DB
	db.Table("Test_User").AutoMigrate(&User{001, "wang"})

	err := db.AutoMigrate(

		system.SysApi{},
		// system.SysUser{},
		// system.SysBaseMenu{},
		// system.JwtBlacklist{},
		// system.SysAuthority{},
		// system.SysDictionary{},
		// system.SysOperationRecord{},
		// system.SysAutoCodeHistory{},
		// system.SysDictionaryDetail{},
		// system.SysBaseMenuParameter{},
		// system.SysBaseMenuBtn{},
		// system.SysAuthorityBtn{},
		// system.SysAutoCode{},
		// system.SysExportTemplate{},

	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
