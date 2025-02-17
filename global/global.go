package global

import (
	"awesomeProject/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CONFIG        config.Server
	VP            *viper.Viper
	DB            *gorm.DB
	ACTIVE_DBNAME *string
	Log           *zap.Logger
)
