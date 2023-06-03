package handler

import (
	"trade/https/trade_api/def"
	"trade/https/trade_api/logic/user"

	"github.com/gin-gonic/gin"
	"github.com/weblazy/easy/http/http_server/service"
)

// 获取用户信息
func GetUserInfo(g *gin.Context) {
	svcCtx := service.NewServiceContext(g)
	req := new(def.GetUserInfoRequest)
	err := svcCtx.BindValidator(req)
	if err != nil {
		svcCtx.Error(err)
		return
	}
	svcCtx.Return(user.GetUserInfo(svcCtx, req))
}
