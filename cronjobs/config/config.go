package config

type Config struct {
	BaseConfig struct{}
}

var Conf = Config{
	BaseConfig: struct{}{},
}

var LocalConfig = ""
