package config

import (
	"github.com/weblazy/easy/http/http_server/http_server_config"
)

type Config struct {
	BaseConfig       struct{}
	HttpServerConfig *http_server_config.Config
}

var Conf = Config{
	BaseConfig:       struct{}{},
	HttpServerConfig: http_server_config.DefaultConfig(),
}

var LocalConfig = ""
