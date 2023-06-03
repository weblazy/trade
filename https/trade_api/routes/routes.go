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

	userGroup := router.Group("/trade")

	userInterceptor(userGroup)
	userGroup.POST("/open_trade", handler.GetUserInfo)
	userGroup.POST("/close_trade", handler.GetUserInfo)
	userGroup.POST("/create_order", handler.GetUserInfo)
	userGroup.POST("/cancel_order", handler.GetUserInfo)

}
