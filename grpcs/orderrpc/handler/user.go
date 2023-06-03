package handler

import (
	"context"

	"trade/grpcs/orderrpc/logic/user_logic"
	"trade/grpcs/orderrpc/proto/user"

	"github.com/weblazy/easy/code_err"
)

type UserService struct {
	user.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

// 获取用户信息
func (h *UserService) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	svcCtx := &user_logic.GetUserInfoCtx{
		Log: code_err.NewLog(ctx),
		Req: req,
		Res: new(user.GetUserInfoResponse),
	}
	err := user_logic.GetUserInfo(svcCtx)
	if err != nil {
		svcCtx.Res.Code = err.Code
		svcCtx.Res.Msg = err.Msg
	}
	return svcCtx.Res, nil
}
