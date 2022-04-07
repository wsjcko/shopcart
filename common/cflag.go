package common

import (
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source"
)

func LoadCflag(cflagSource source.Source) (config.Config, error) {
	config, err := config.NewConfig()
	if err != nil {
		return config, err
	}
	//加载配置
	err = config.Load(cflagSource)
	return config, err
}
