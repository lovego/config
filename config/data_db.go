package config

import (
	"fmt"
	"sort"

	"github.com/lovego/strmap"
)

func GetDB(configMap strmap.StrMap, typ, key string) (interface{}, error) {
	v, ok := configMap.Get(typ)[key]
	if !ok {
		return nil, fmt.Errorf("db config `%s.%s` not found.", typ, key)
	}
	switch value := v.(type) {
	case string:
		return value, nil
	case map[interface{}]interface{}:
		return GetShards(value, typ+`.`+key)
	default:
		return nil, fmt.Errorf(
			"db config `%s.%s` should be a string or map, but got: %v", typ, key, v,
		)
	}
}

type Shards struct {
	Shards   []Shard
	Settings ShardsSettings
}

type Shard struct {
	No  int
	Url string
}

type ShardsSettings struct {
	IdSeqIncrementBy int
}

func GetShards(m map[interface{}]interface{}, path string) (*Shards, error) {
	var shardsConfig Shards
	for k, v := range m {
		if k == "settings" {
			if settings, err := GetShardsSettings(v, path); err != nil {
				return nil, err
			} else {
				shardsConfig.Settings = *settings
			}
		} else if shardNo, ok := k.(int); ok {
			if shardUrl, ok := v.(string); ok {
				shardsConfig.Shards = append(shardsConfig.Shards, Shard{shardNo, shardUrl})
			} else {
				return nil, fmt.Errorf("`%s.%d` should be a string, but got: %v", path, k, v)
			}
		} else {
			return nil, fmt.Errorf(
				"`%s` invalid shard number: %v, it should be an integer.", path, k,
			)
		}
	}
	sort.Slice(shardsConfig.Shards, func(i, j int) bool {
		return shardsConfig.Shards[i].No < shardsConfig.Shards[j].No
	})

	return &shardsConfig, nil
}

func GetShardsSettings(v interface{}, path string) (*ShardsSettings, error) {
	//m, ok := v.(map[string]interface{})
	//if !ok {
	//	return nil, fmt.Errorf("`%s.settings` should be a map, but got: %v", path, v)
	//}

	var val map[string]interface{}
	switch v1 := v.(type) {
	case strmap.StrMap:
		val = (map[string]interface{})(v1)
	case map[interface{}]interface{}:
		val = v.(map[string]interface{})
	case map[string]interface{}:
		val = v.(map[string]interface{})
	default:
		return nil, fmt.Errorf("`%s.settings` should be a map, but got: %v", path, v)
	}

	//m = val

	var settings ShardsSettings

	for k, v := range val {
		switch k {
		case "idSeqIncrementBy":
			if i, ok := v.(int); ok {
				settings.IdSeqIncrementBy = i
			}
		default:
			return nil, fmt.Errorf("`%s.settings` unexpected key: %v", path, k)
		}
	}
	return &settings, nil
}
