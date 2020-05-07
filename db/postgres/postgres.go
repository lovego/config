package postgres

import (
	"database/sql"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
	"github.com/lovego/bsql"
	"github.com/lovego/config"
	"github.com/lovego/config/db/dburl"
)

var dbs = struct {
	sync.Mutex
	m map[string]*bsql.DB
}{m: make(map[string]*bsql.DB)}

func DB(name string) *bsql.DB {
	return Get(config.Get("postgres").GetString(name))
}

func Get(dbAddr string) *bsql.DB {
	dbs.Lock()
	defer dbs.Unlock()

	db := dbs.m[dbAddr]
	if db == nil {
		db = New(dbAddr)
		dbs.m[dbAddr] = db
	}
	return db
}

func New(dbAddr string) *bsql.DB {
	return NewWithTimeout(dbAddr, 5*time.Second)
}

func NewWithTimeout(dbAddr string, timeout time.Duration) *bsql.DB {
	dbUrl := dburl.Parse(dbAddr)
	db, err := sql.Open("postgres", dbUrl.URL.String())
	if err != nil {
		log.Panic(err)
	}
	if err := db.Ping(); err != nil {
		log.Panic(err)
	}
	db.SetMaxOpenConns(dbUrl.MaxOpen)
	db.SetMaxIdleConns(dbUrl.MaxIdle)
	db.SetConnMaxLifetime(dbUrl.MaxLife)
	return bsql.New(db, timeout)
}
