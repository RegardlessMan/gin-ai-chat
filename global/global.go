package global

import (
	"awesomeProject/config"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	CONFIG        config.Server
	VP            *viper.Viper
	DB            *gorm.DB
	ACTIVE_DBNAME *string
)
