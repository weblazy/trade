package orderrpc

import (
	"log"

	"github.com/urfave/cli/v2"

	"trade/grpcs/orderrpc/config"
	"trade/grpcs/orderrpc/handler"
	"trade/grpcs/orderrpc/proto/user"

	"github.com/weblazy/easy/closes"
	"github.com/weblazy/easy/econfig"
	"github.com/weblazy/easy/grpc/grpc_server"
)

var Cmd = &cli.Command{
	Name:    "orderrpc",
	Aliases: []string{},
	Usage:   "orderrpc start",
	Subcommands: []*cli.Command{
		{
			Name:   "start",
			Usage:  "start service",
			Action: Run,
		},
	},
}

func Run(c *cli.Context) error {
	defer closes.Close()
	econfig.InitGlobalViper(&config.Conf, config.LocalConfig)
	s := grpc_server.NewGrpcServer(config.Conf.GrpcServerConfig)
	userService := handler.NewUserService()

	user.RegisterUserServiceServer(s, userService)
	err := s.Init()
	if err != nil {
		log.Fatal(err)
	}
	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
