package config

import (
	"fmt"
	"log"
	"net/url"
	"sort"
	"strconv"

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
	case map[string]interface{}:
		return GetShards2(value, typ+`.`+key)
	case strmap.StrMap:
		return GetShards2(map[string]interface{}(value), typ+`.`+key)
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
	DbAddr string
	No     int
	Url    string
}

type ShardsSettings struct {
	IdSeqIncrementBy int
}

func GetShards(m map[interface{}]interface{}, path string) (*Shards, error) {
	var shards = &Shards{}
	for k, v := range m {
		if err := parseShard(shards, k, v, path); err != nil {
			return nil, err
		}
	}
	sort.Slice(shards.Shards, func(i, j int) bool {
		return shards.Shards[i].No < shards.Shards[j].No
	})
	return shards, nil
}

func GetShards2(m map[string]interface{}, path string) (*Shards, error) {
	var shards = &Shards{}
	for k, v := range m {
		if err := parseShard(shards, k, v, path); err != nil {
			return nil, err
		}
	}
	sort.Slice(shards.Shards, func(i, j int) bool {
		return shards.Shards[i].No < shards.Shards[j].No
	})
	return shards, nil
}

func parseShard(shards *Shards, k, v interface{}, path string) error {
	if k == "settings" {
		if settings, err := GetShardsSettings(v, path); err != nil {
			return err
		} else {
			shards.Settings = settings
			return nil
		}
	}

	var shardNo int
	switch key := k.(type) {
	case string:
		if i, err := strconv.Atoi(key); err != nil {
			return fmt.Errorf("`%s` invalid shard number: %v, it should be an integer.", path, k)
		} else {
			shardNo = i
		}
	case int:
		shardNo = key
	default:
		return fmt.Errorf("`%s` invalid shard number: %v, it should be an integer.", path, k)
	}

	if shardUrl, ok := v.(string); ok {
		u, err := url.Parse(shardUrl)
		if err != nil {
			log.Panic(err)
		}
		shards.Shards = append(shards.Shards, Shard{
			DbAddr: u.Host,
			No:     shardNo,
			Url:    shardUrl,
		})
	} else {
		return fmt.Errorf("`%s.%d` should be a string, but got: %v", path, k, v)
	}
	return nil
}

func GetShardsSettings(v interface{}, path string) (ShardsSettings, error) {
	var val map[string]interface{}
	switch v1 := v.(type) {
	case strmap.StrMap:
		val = v1
	case map[interface{}]interface{}:
		val = v.(map[string]interface{})
	case map[string]interface{}:
		val = v.(map[string]interface{})
	default:
		return ShardsSettings{}, fmt.Errorf("`%s.settings` should be a map, but got: %v", path, v)
	}

	var settings ShardsSettings

	for k, v := range val {
		switch k {
		case "idSeqIncrementBy":
			if i, ok := v.(int); ok {
				settings.IdSeqIncrementBy = i
			}
		default:
			return ShardsSettings{}, fmt.Errorf("`%s.settings` unexpected key: %v", path, k)
		}
	}
	return settings, nil
}
