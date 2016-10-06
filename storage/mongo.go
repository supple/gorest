package storage

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

func NewMongo(url string, dbName string) *Mongo {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(fmt.Sprintf("Mongodb not found %s", url))
	}
	m := &Mongo{session: session, dbName: dbName}

	return m
}

type Mongo struct {
	session *mgo.Session
	// db mgo.Database
	dbName  string
}

func (m *Mongo) Coll(collectionName string) *mgo.Collection {
	return m.session.DB(m.dbName).C(collectionName)
}