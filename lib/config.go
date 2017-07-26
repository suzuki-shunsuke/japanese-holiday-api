package lib

import (
	"github.com/BurntSushi/toml"
	"github.com/suzuki-shunsuke/japanese-holiday-api/types"
)

func GetConfig() (*types.Config, error) {
	var config types.Config
	_, err := toml.DecodeFile("config.toml", &config)
	return &config, err
}
