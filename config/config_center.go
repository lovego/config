package config

import (
	"log"

	"github.com/lovego/config_sdk/go_config_sdk"
	"gopkg.in/yaml.v3"
)


type ConfigCenter struct {
	Secret  string `yaml:"secret" json:"secret" c:"访问密码"`
	Pull    string `yaml:"pull" json:"pull" c:"访问地址"`
	Project string `yaml:"project" json:"project" c:"项目"`
	Version string `yaml:"version" json:"version" c:"版本"`
}

func GetCenterConfig(center ConfigCenter, env string) *Config {
	arg := go_config_sdk.ConfigTag{
		Project:      center.Project,
		Env:          env,
		EndPointType: "server",
		Version:      center.Version,
		Hash:         "",
	}
	config, err := go_config_sdk.GetConfig(center.Pull, center.Secret, arg)
	if err != nil {
		log.Fatalf(`config center:%s`, err.Error())
	}

	data, err := TranRemoteToLocal(config)
	if err != nil {
		log.Fatalf(`config center:%s`, err.Error())
	}
	return data
}

func TranRemoteToLocal(config *go_config_sdk.Config) (*Config, error) {

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
