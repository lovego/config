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
var rootDir = getRoot(theEnv)
var configDir = filepath.Join(Root(), "config")

func Env() *config.Environment {
	return &theEnv
}
func Root() string {
	return rootDir
}
func Dir() string {
	return configDir
}
func FilePath(env *config.Environment) string {
	return filepath.Join(Dir(), env.Minor()+`.yml`)
}

func getEnv() string {
	env := os.Getenv(config.EnvVar)
	if env == "" {
		if strings.HasSuffix(os.Args[0], ".test") {
			env = "test"
		} else {
			env = "dev"
		}
	}
	return env
}

func getRoot(env config.Environment) string {
	root := detectRootByExecutable(env.Minor())
	if root == "" {
		if release := config.DetectReleaseConfigDirOf(env.Major()); release != "" {
			root = filepath.Join(release, "img-app")
		}
	}
	if root == "" {
		log.Panic("app root not found.")
	}
	return root
}

func detectRootByExecutable(minorEnv string) string {
	executable, err := filepath.Abs(os.Args[0])
	if err != nil {
		log.Panic(err)
	}
	executableDir := filepath.Dir(executable)
	if fs.Exist(filepath.Join(executableDir, "config", minorEnv+".yml")) {
		return executableDir
	}
	return ""
}
