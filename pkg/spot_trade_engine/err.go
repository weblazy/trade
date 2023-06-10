package spot_trade_engine

import "github.com/weblazy/easy/code_err"

var (
	WorkerNotFoundErr = code_err.NewCodeErr(110020, "验证码发送失败")
	TradeClosedErr    = code_err.NewCodeErr(110020, "验证码发送失败")
	OrderChanCloseErr = code_err.NewCodeErr(110020, "验证码发送失败")
	OrderNotFoundErr  = code_err.NewCodeErr(110020, "验证码发送失败")
	OrderExistErr     = code_err.NewCodeErr(110020, "验证码发送失败")
)
