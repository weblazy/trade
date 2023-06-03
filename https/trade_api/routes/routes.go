package routes

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"

	"trade/https/trade_api/handler"
)

func Routes(router *gin.Engine) {

	// 根目录健康检查
	router.Any("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Easy Service")
	})

	userGroup := router.Group("/app/user")

	userInterceptor(userGroup)
	userGroup.POST("/user/get_user_info", handler.GetUserInfo) //获取用户信息

}
