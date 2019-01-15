package mysql

import (
	"database/sql"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lovego/bsql"
	"github.com/lovego/config"
	"github.com/lovego/config/db/dburl"
)

var postgresDBs = struct {
	sync.Mutex
	m map[string]*bsql.DB
}{m: make(map[string]*bsql.DB)}

func DB(name string) *bsql.DB {
	postgresDBs.Lock()
	defer postgresDBs.Unlock()
	db := postgresDBs.m[name]
	if db == nil {
		db = NewDB(config.Get("mysql").GetString(name))
		postgresDBs.m[name] = db
	}
	return db
}

func NewDB(dbAddr string) *bsql.DB {
	dbUrl := dburl.Parse(dbAddr)
	db, err := sql.Open("mysql", dbUrl.URL.String())
	if err != nil {
		log.Panic(err)
	}
	if err := db.Ping(); err != nil {
		log.Panic(err)
	}
	db.SetMaxOpenConns(dbUrl.MaxOpen)
	db.SetMaxIdleConns(dbUrl.MaxIdle)
	db.SetConnMaxLifetime(dbUrl.MaxLife)
	return bsql.New(db, 5*time.Second)
}
