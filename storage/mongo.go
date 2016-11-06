package storage

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

var instances map[string]*MongoDB = make(map[string]*MongoDB)

func init()  {
    db := NewMongoDB("192.168.1.106:27017", "lcache")
    instances["entities"] = db

	dbEvents := NewMongoDB("192.168.1.106:27017", "events")
	instances["events"] = dbEvents
}

func SetInstance(name string, db *MongoDB) {
    instances[name] = db
}

func GetInstance(name string) *MongoDB {
	return instances[name]
}

func NewMongoDB(url string, dbName string) *MongoDB {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(fmt.Sprintf("Mongodb not found %s", url))
	}
	m := &MongoDB{session: session, dbName: dbName}

	return m
}

type MongoDB struct {
	session *mgo.Session
	// db mgo.Database
	dbName  string
}

func (m *MongoDB) Coll(collName string) *mgo.Collection {
	return m.session.DB(m.dbName).C(collName)
}

func DropCollection(db *MongoDB, cn string) {
    db.Coll(cn).DropCollection()
}