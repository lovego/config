package config

import (
	"log"
	"path/filepath"

	"github.com/lovego/config/config"
	"github.com/lovego/strmap"
)

var theData = config.Data(filepath.Join(Dir(), `envs/`+Env().Minor()+`.yml`))

func Get(key string) strmap.StrMap {
	return theData.Get(key)
}

func GetString(key string) string {
	return theData.GetString(key)
}

func GetSlice(key string) []strmap.StrMap {
	return theData.GetSlice(key)
}

func GetStringSlice(key string) []string {
	return theData.GetStringSlice(key)
}

func GetDBConfig(typ, key string) interface{} {
	v, err := config.GetDB(theData, typ, key)
	if err != nil {
		log.Panic(v)
	}
	return v
}
