package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/lovego/config/config"
	"github.com/lovego/fs"
)

var theEnv = config.NewEnv(getEnv())
var rootDir = getRoot()
var configDir = filepath.Join(Root(), Env().ConfigDir())

func Env() *config.Environment {
	return &theEnv
}
func Root() string {
	return rootDir
}
func Dir() string {
	return configDir
}

func getEnv() string {
	env := os.Getenv(config.EnvVar)
	if env == `` {
		if strings.HasSuffix(os.Args[0], `.test`) {
			env = `test`
		} else {
			env = `dev`
		}
	}
	return env
}

func getRoot() string {
	program, err := filepath.Abs(os.Args[0])
	if err != nil {
		log.Panic(err)
	}
	configYML := Env().ConfigDir() + "/config.yml"
	if programDir := filepath.Dir(program); fs.Exist(filepath.Join(programDir, configYML)) {
		return programDir
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	projectDir := fs.DetectDir(cwd, `release/img-app/`+configYML)
	if projectDir == `` {
		log.Panic(`app root not found.`)
	}
	return filepath.Join(projectDir, `release/img-app`)
}
