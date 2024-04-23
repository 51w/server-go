package initialize

import (
	"github.com/51w/server-go/server/global"
	// "github.com/51w/server-go/server/middleware"
	"github.com/51w/server-go/server/router"
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

	// 设置为发布模式
	if global.GVA_CONFIG.System.Env == "public" {
		gin.SetMode(gin.ReleaseMode) //DebugMode ReleaseMode TestMode
	}

	Router := gin.New()
	Router.Use(gin.Recovery())
	if global.GVA_CONFIG.System.Env != "public" {
		Router.Use(gin.Logger())
	}

	// InstallPlugin(Router) // 安装插件
	systemRouter := router.RouterGroupApp.System
	// exampleRouter := router.RouterGroupApp.Example

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		// systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	// PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	// PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	// {
	// 	systemRouter.InitApiRouter(PrivateGroup, PublicGroup)       // 注册功能api路由
	// 	systemRouter.InitJwtRouter(PrivateGroup)                    // jwt相关路由
	// 	systemRouter.InitUserRouter(PrivateGroup)                   // 注册用户路由
	// 	systemRouter.InitMenuRouter(PrivateGroup)                   // 注册menu路由
	// 	systemRouter.InitSystemRouter(PrivateGroup)                 // system相关路由
	// 	systemRouter.InitCasbinRouter(PrivateGroup)                 // 权限相关路由
	// 	systemRouter.InitAutoCodeRouter(PrivateGroup)               // 创建自动化代码
	// 	systemRouter.InitAuthorityRouter(PrivateGroup)              // 注册角色路由
	// 	systemRouter.InitSysDictionaryRouter(PrivateGroup)          // 字典管理
	// 	systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)        // 自动化代码历史
	// 	systemRouter.InitSysOperationRecordRouter(PrivateGroup)     // 操作记录
	// 	systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)    // 字典详情管理
	// 	systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)     // 字典详情管理
	// 	systemRouter.InitSysExportTemplateRouter(PrivateGroup)      // 导出模板
	// 	exampleRouter.InitCustomerRouter(PrivateGroup)              // 客户路由
	// 	exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由

	// }

	global.GVA_LOG.Info("router register success")
	return Router
}
