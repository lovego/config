package mongo

import (
	"log"
	"sync"

	"github.com/lovego/config"
	"gopkg.in/mgo.v2"
)

var dbs = struct {
	m map[string]Sess
	sync.Mutex
}{
	m: make(map[string]Sess),
}

func Session(name string) Sess {
	return Get(config.Get(`mongo`).GetString(name))
}

func Get(dbAddr string) Sess {
	dbs.Lock()
	defer dbs.Unlock()

	db, ok := dbs.m[dbAddr]
	if !ok {
		db = New(dbAddr)
		dbs.m[dbAddr] = db
	}
	return db
}

func New(dbAddr string) Sess {
	session, err := mgo.Dial(dbAddr)
	if err != nil {
		log.Panic(err)
	}
	return Sess{session}
}

type Sess struct {
	s *mgo.Session
}
type DB struct {
	db *mgo.Database
}
type Coll struct {
	c *mgo.Collection
}

func (s Sess) Session(work func(*mgo.Session)) {
	sess := s.s.Copy()
	defer sess.Close()
	work(sess)
}

func (s Sess) DB(name string) DB {
	return DB{s.s.DB(name)}
}

func (db DB) Session(work func(*mgo.Database)) {
	sess := db.db.Session.Copy()
	defer sess.Close()
	work(db.db.With(sess))
}

func (db DB) C(name string) Coll {
	return Coll{db.db.C(name)}
}

func (c Coll) Session(work func(*mgo.Collection)) {
	sess := c.c.Database.Session.Copy()
	defer sess.Close()
	work(c.c.With(sess))
}
