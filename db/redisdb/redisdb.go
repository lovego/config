package redisdb

import (
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/lovego/config"
)

var redisPools = struct {
	sync.Mutex
	m map[string]*redis.Pool
}{m: make(map[string]*redis.Pool)}

func Pool(name string) *redis.Pool {
	redisPools.Lock()
	defer redisPools.Unlock()
	pool := redisPools.m[name]
	if pool == nil {
		pool = newPool(name)
		redisPools.m[name] = pool
	}
	return pool
}

func newPool(name string) *redis.Pool {
	var url = config.Get(`redis`).GetString(name)
	return &redis.Pool{
		MaxIdle:     32,
		MaxActive:   128,
		IdleTimeout: 600 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(
				url,
				redis.DialConnectTimeout(3*time.Second),
				redis.DialReadTimeout(3*time.Second),
				redis.DialWriteTimeout(3*time.Second),
			)
		},
	}
}

func Do(name string, work func(redis.Conn)) {
	conn := Pool(name).Get()
	defer conn.Close()
	work(conn)
}

func SubscribeConn(name string) (redis.Conn, error) {
	return redis.DialURL(
		config.Get(`redis`).GetString(name),
		redis.DialConnectTimeout(3*time.Second),
		redis.DialWriteTimeout(3*time.Second),
	)
}
