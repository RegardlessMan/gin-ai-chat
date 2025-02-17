package internal

import (
	"awesomeProject/config"
	"fmt"
	"gorm.io/gorm/logger"
)

type Writer struct {
	config config.GeneralDB
	writer logger.Writer
}

func NewWriter(config config.GeneralDB) *Writer {
	return &Writer{config: config}
}

// Printf 格式化打印日志
func (c *Writer) Printf(message string, data ...any) {

	// 当有日志时候均需要输出到控制台
	fmt.Printf(message, data...)

	// 当开启了zap的情况，会打印到日志记录
	//if c.config.LogZap {
	//	switch c.config.LogLevel() {
	//	case logger.Silent:
	//		global.Log.Debug(fmt.Sprintf(message, data...))
	//	case logger.Error:
	//		global.Log.Error(fmt.Sprintf(message, data...))
	//	case logger.Warn:
	//		global.Log.Warn(fmt.Sprintf(message, data...))
	//	case logger.Info:
	//		global.Log.Info(fmt.Sprintf(message, data...))
	//	default:
	//		global.Log.Info(fmt.Sprintf(message, data...))
	//	}
	//	return
	//}
}
