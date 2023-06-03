package main

import (
	"context"
	"os"
	"trade/common"
	"trade/cronjobs"
	"trade/grpcs/orderrpc"
	"trade/https/trade_api"
	"trade/jobs"

	"github.com/sunmi-OS/gocore/v2/utils"
	"github.com/urfave/cli/v2"
	"github.com/weblazy/easy/elog"
)

func main() {
	// 打印Banner
	utils.PrintBanner(common.ProjectName)
	// 配置cli参数
	cliApp := cli.NewApp()
	cliApp.Name = common.ProjectName
	cliApp.Version = common.ProjectVersion

	// 指定命令运行的函数
	cliApp.Commands = []*cli.Command{
		trade_api.Cmd, orderrpc.Cmd, cronjobs.Cmd, jobs.Cmd,
	}

	// 启动cli
	if err := cliApp.Run(os.Args); err != nil {
		elog.ErrorCtx(context.Background(), "Failed to start application", elog.FieldError(err))
	}
}
