package initialize

import (
	"awesomeProject/global"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.CONFIG.System.DbType {
	case "mysql":
		global.ACTIVE_DBNAME = &global.CONFIG.Mysql.Dbname
		return GormMysql()
	default:
		global.ACTIVE_DBNAME = &global.CONFIG.Mysql.Dbname
		return GormMysql()
	}
}
