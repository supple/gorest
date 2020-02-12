package storage

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

var instances map[string]*MongoDB = make(map[string]*MongoDB)

func SetInstance(name string, db *MongoDB) {
	instances[name] = db
}

func GetInstance(name string) *MongoDB {
	if val, ok := instances[name]; ok {
		return val
	}
	panic(fmt.Sprintf("Database `%s` not defined", name))
}

func NewMongoDB(url string, dbName string) *MongoDB {
	session, err := mgo.DialWithTimeout(url, time.Second*5)
	if err != nil {
		panic(fmt.Sprintf("Mongodb not found %s", url))
	}

	m := &MongoDB{Session: session, dbName: dbName}

	return m
}

type MongoDB struct {
	Session *mgo.Session
	// db mgo.Database
	dbName string
}

func (m *MongoDB) Coll(collName string) *mgo.Collection {
	//m.Session.Refresh()
	return m.Session.DB(m.dbName).C(collName)
}

func DropCollection(db *MongoDB, cn string) {
	db.Coll(cn).DropCollection()
}

func DropDatabase(db *MongoDB) {
	db.Session.DB(db.dbName).DropDatabase()
}
