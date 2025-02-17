package core

import (
	"awesomeProject/core/internal"
	"awesomeProject/global"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
// Author [SliverHorn](https://github.com/SliverHorn)
func Viper(path ...string) *viper.Viper {

	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" {
			// 命令行为空，判断环境变量
			if configEnv := os.Getenv("CONFIG_PATH"); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
				case gin.TestMode:
					config = internal.ConfigTestFile
				}
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.Mode(), config)
			} else {

			}

		} else {
			fmt.Printf("您正在使用命令行-c传递的至，config路径为%s\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig() //监听配置文件变化
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
