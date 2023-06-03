package user_logic

import (
	"trade/grpcs/orderrpc/proto/user"

	"github.com/weblazy/easy/code_err"
)

type GetUserInfoCtx struct {
	*code_err.Log
	Req *user.GetUserInfoRequest
	Res *user.GetUserInfoResponse
}

// 获取用户信息
func GetUserInfo(ctx *GetUserInfoCtx) *code_err.CodeErr {
	return nil
}
