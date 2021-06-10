package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Name string
	Envs map[string]*EnvConfig
	path string
}

func Get(path, envMajor string) *Config {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	config := &Config{path: path}
	if err := yaml.Unmarshal(content, config); err != nil {
		log.Fatalf("parse %s: %v", path, err)
	}

	if envMajor != "" {
		envs := make(map[string]*EnvConfig, len(config.Envs))
		for envMinor, envConfig := range config.Envs {
			envs[envMajor+"."+envMinor] = envConfig
		}
		config.Envs = envs
	}

	for env, envConfig := range config.Envs {
		if envConfig != nil {
			envConfig.init(config.Name, env)
		}
	}
	return config
}

func (config *Config) Get(env string) *EnvConfig {
	envConfig := config.Envs[env]
	if envConfig == nil {
		e := NewEnv(env)
		log.Fatalf("%s envs.%s: not defined.", config.path, e.Minor())
	}
	return envConfig
}
