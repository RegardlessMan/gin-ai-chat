package initialize

import (
	"awesomeProject/global"
	"github.com/jinzhu/gorm"
)

func Gorm() *gorm.DB {
	switch global.CONFIG.System.DbType {
	case "mysql":
		global.ACTIVE_DBNAME = &global.CONFIG.Mysql.Dbname
	}
	return nil
}
