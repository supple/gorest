package storage

import (
	"gopkg.in/mgo.v2"
	"fmt"
    "time"
)

var instances map[string]*MongoDB = make(map[string]*MongoDB)

func init()  {
    //db := NewMongoDB("192.168.1.106:27017", "lcache")
    //instances["entities"] = db
    //
	//dbEvents := NewMongoDB("192.168.1.106:27017", "events")
	//instances["events"] = dbEvents
}

func SetInstance(name string, db *MongoDB) {
    instances[name] = db
}

func GetInstance(name string) *MongoDB {
	return instances[name]
}

func NewMongoDB(url string, dbName string) *MongoDB {
	session, err := mgo.DialWithTimeout(url, time.Second * 5)
	if err != nil {
		panic(fmt.Sprintf("Mongodb not found %s", url))
	}

	m := &MongoDB{Session: session, dbName: dbName}

	return m
}

type MongoDB struct {
    Session *mgo.Session
	// db mgo.Database
    dbName  string
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