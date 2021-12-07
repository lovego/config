package config

import (
	"log"

	"github.com/lovego/config_sdk/go_config_sdk"
	"gopkg.in/yaml.v3"
)

func GetCenterConfig(center ConfigCenter, env string) *Config {
	arg := go_config_sdk.Arg{
		Project:      center.Project,
		Env:          env,
		EndPointType: "server",
		Version:      center.Version,
		Hash:         "",
	}
	config, err := go_config_sdk.GetConfig(center.Addr, center.Secret, arg)
	if err != nil {
		log.Fatalf(`config center:%s`, err.Error())
	}

	data, err := TranCenterToLocal(config)
	if err != nil {
		log.Fatalf(`config center:%s`, err.Error())
	}
	return data
}

func TranCenterToLocal(config *go_config_sdk.Config) (*Config, error) {

	content, err := yaml.Marshal(config)
	if err != nil {
		return nil, err
	}

	data := new(Config)
	err = yaml.Unmarshal(content, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
