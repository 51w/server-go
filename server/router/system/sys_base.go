package system

import (
	v1 "github.com/51w/server-go/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.GET("login", baseApi.Login)
		// baseRouter.POST("captcha", baseApi.Captcha)
	}
	return baseRouter
}
