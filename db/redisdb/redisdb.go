package redisdb

import (
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/lovego/config"
)

var dbs = struct {
	sync.Mutex
	m map[string]*redis.Pool
}{m: make(map[string]*redis.Pool)}

func Pool(name string) *redis.Pool {
	return Get(config.Get(`redis`).GetString(name))
}

func Get(dbAddr string) *redis.Pool {
	dbs.Lock()
	defer dbs.Unlock()

	db := dbs.m[dbAddr]
	if db == nil {
		db = New(dbAddr)
		dbs.m[dbAddr] = db
	}
	return db
}

func New(dbUrl string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     32,
		MaxActive:   128,
		IdleTimeout: 600 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(
				dbUrl,
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
