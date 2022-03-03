package config

import (
	"log"

	"github.com/lovego/config/config"
	"github.com/lovego/strmap"
)

func Get(key string) strmap.StrMap {
	return theConfig.Data.Get(key)
}

func GetString(key string) string {
	return theConfig.Data.GetString(key)
}

func GetSlice(key string) []strmap.StrMap {
	return theConfig.Data.GetSlice(key)
}

func GetStringSlice(key string) []string {
	return theConfig.Data.GetStringSlice(key)
}

func GetDBConfig(typ, key string) interface{} {
	v, err := config.GetDB(theConfig.Data, typ, key)
	if err != nil {
		log.Panic(err)
	}
	return v
}
