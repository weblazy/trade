package jobs

import (
	"trade/jobs/config"
	"trade/jobs/handler"

	"github.com/urfave/cli/v2"
	"github.com/weblazy/easy/closes"
	"github.com/weblazy/easy/econfig"
)

// Job cmd 任务相关
var Cmd = &cli.Command{
	Name:    "job",
	Aliases: []string{"j"},
	Usage:   "job",
	Subcommands: []*cli.Command{
		{
			Name:   "InitUser",
			Usage:  "初始化默认用户",
			Action: InitUser,
		},
	},
}

func InitUser(c *cli.Context) error {
	defer closes.Close()
	// 初始化必要内容
	econfig.InitGlobalViper(&config.Conf, config.LocalConfig)
	handler.InitUser()
	return nil
}
