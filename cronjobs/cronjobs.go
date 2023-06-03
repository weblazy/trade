package cronjobs

import (
	"trade/cronjobs/config"
	"trade/cronjobs/handler"

	"github.com/robfig/cron/v3"
	"github.com/urfave/cli/v2"
	"github.com/weblazy/easy/closes"
	"github.com/weblazy/easy/econfig"
)

var Cmd = &cli.Command{
	Name:    "cron",
	Aliases: []string{"c"},
	Usage:   "cron start",
	Subcommands: []*cli.Command{
		{
			Name:   "start",
			Usage:  "开启运行api服务",
			Action: Run,
		},
	},
}

func Run(c *cli.Context) error {
	defer closes.Close()
	// 初始化必要内容
	econfig.InitGlobalViper(&config.Conf, config.LocalConfig)
	cronJob := cron.New()

	_, _ = cronJob.AddFunc("@every 30m", handler.SyncUser)

	cronJob.Start()

	closes.AddShutdown(closes.ModuleClose{
		Name:     "CronTable",
		Priority: 0,
		Func: func() {
			cronJob.Stop()
		},
	})
	closes.SignalClose()
	return nil
}
