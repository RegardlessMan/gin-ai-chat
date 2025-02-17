package initialize

import (
	"awesomeProject/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	//systemRouter := router.RouterGroupApp.System
	//exampleRouter := router.RouterGroupApp.Example
	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面3行注释
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/assets", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	Router.StaticFS(global.CONFIG.Local.StorePath, justFilesFilesystem{http.Dir(global.CONFIG.Local.StorePath)}) // Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	// global.GVA_LOG.Info("use middleware cors")

	//接入swagger
	//docs.SwaggerInfo.BasePath = global.CONFIG.System.RouterPrefix
	//Router.GET(global.CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//global.Log.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	PrivateGroup := Router.Group(global.CONFIG.System.RouterPrefix)

	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
		// 健康监测
		PrivateGroup.GET("/private/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	//插件路由安装
	//InstallPlugin(PrivateGroup, PublicGroup, Router)

	// 注册业务路由
	//initBizRouter(PrivateGroup, PublicGroup)

	global.GROUTERS = Router.Routes()

	global.Log.Info("router register success")
	return Router
}
