package trade_api

import (
	"trade/https/trade_api/config"
	"trade/https/trade_api/routes"

	"github.com/urfave/cli/v2"
	"github.com/weblazy/easy/closes"
	"github.com/weblazy/easy/econfig"
	"github.com/weblazy/easy/http/http_server"
)

var Cmd = &cli.Command{
	Name:    "api",
	Aliases: []string{"a"},
	Usage:   "api start",
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
	econfig.InitGlobalViper(&config.Conf, config.LocalConfig)

	s, err := http_server.NewHttpServer(config.Conf.HttpServerConfig)
	if err != nil {
		return err
	}
	// 注册路由
	routes.Routes(s.Engine)

	err = s.Start()
	if err != nil {
		return err
	}
	return nil
}
