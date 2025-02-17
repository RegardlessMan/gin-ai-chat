package global

import (
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
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
	GROUTERS      gin.RoutesInfo
)
