package core

import (
    "gopkg.in/mgo.v2/bson"
    "github.com/supple/gorest/storage"
    "reflect"
    "gopkg.in/mgo.v2"
    "io"
    "fmt"
)

const CUSTOMER_NAME_FIELD string = "customerName"

type Adapter func(query bson.M) bson.M

func ContextDecorator(cc *CustomerContext) Adapter {
    return func (query bson.M) bson.M {
        query[CUSTOMER_NAME_FIELD] = cc.CustomerName
        return query
    }
}

func EmptyDecorator() Adapter {
    return func (query bson.M) bson.M {
        return query
    }
}

// Gateway

type Gateway struct {
    collectionName string
    cc *CustomerContext
    coll *mgo.Collection
    db *storage.MongoDB
    decorate Adapter
}

func NewGateway(collectionName string, decorate Adapter, db *storage.MongoDB) *Gateway {
    return &Gateway{collectionName: collectionName, db: db, decorate: decorate}
}

func (gt *Gateway) Insert(model interface{}) error {
    // generate id
    vDst := reflect.Indirect(reflect.ValueOf(model)).FieldByName("Id")
    if (vDst.Len() == 0) {
        vDst.SetString(NewId())
    }

    coll := gt.db.Coll(gt.collectionName)
    err := coll.Insert(model)
    logQuery(coll, model)
    handleError(gt.db, err)

    return gt.toApiError(err)
}

//func (gt *Gateway) Update(id string, model *map[string]interface{}) error {
func (gt *Gateway) Update(id string, model interface{}) error {
    query := gt.decorate(bson.M{"_id": id})
    coll := gt.db.Coll(gt.collectionName)
    err := coll.Update(query, model)
    logQuery(coll, struct {query interface{} ; model interface{}} {query, model})
    handleError(gt.db, err)

    return gt.toApiError(err)
}

func (gt *Gateway) Remove(id string) error {
    query := gt.decorate(bson.M{"_id": id})
    coll := gt.db.Coll(gt.collectionName)
    err := coll.Remove(query)
    logQuery(coll, query)
    handleError(gt.db, err)

    return gt.toApiError(err)
}

func (gt *Gateway) FindById(id string, result interface{}) error  {
    query := gt.decorate(bson.M{"_id": id})
    coll := gt.db.Coll(gt.collectionName)
    err := coll.Find(query).One(result)
    logQuery(coll, query)
    handleError(gt.db, err)

    return gt.toApiError(err)
}

func (gt *Gateway) FindOneBy(query bson.M, result interface{}) error  {
    query = gt.decorate(query)
    coll := gt.db.Coll(gt.collectionName)
    err := coll.Find(query).One(result)
    logQuery(coll, query)
    handleError(gt.db, err)

    return gt.toApiError(err)
}

func (gt *Gateway) toApiError(err error) error {
    if (err == io.EOF) {
        return ErrDatabase
    }
    if (err == mgo.ErrNotFound) {
        //return mgo.ErrNotFound
        return ErrNotFound
        //return &ErrObjectNotFound{gt.collectionName, ""}
    }

    return nil
}

func logQuery(coll *mgo.Collection, query interface{})  {
    Log(fmt.Sprintf("[query] %s, %s", coll.FullName, query))
}

func handleError(db *storage.MongoDB, err error) {
    if (err == io.EOF) {
        db.Session.Refresh()
    }
}