package core

import (
    "gopkg.in/mgo.v2/bson"
    s "github.com/supple/gorest/storage"
    lc "github.com/supple/gorest/utils"
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
    db *s.MongoDB
    decorate Adapter
}

func NewGateway(collectionName string, decorate Adapter, db *s.MongoDB) *Gateway {
    return &Gateway{collectionName: collectionName, db: db, decorate: decorate}
}

func (gt *Gateway) Insert(model interface{}) error {
    // generate id
    vDst := reflect.Indirect(reflect.ValueOf(model)).FieldByName("Id")
    if (vDst.Len() == 0) {
        vDst.SetString(lc.NewId())
    }

    err := gt.db.Coll(gt.collectionName).Insert(model)
    handleError(gt.db, err)

    return gt.toApiError(err)
}

func (gt *Gateway) Update(id string, model *map[string]interface{}) error {
    query := gt.decorate(bson.M{"_id": id})
    err := gt.db.Coll(gt.collectionName).Update(query, model)
    handleError(gt.db, err)

    return gt.toApiError(err)
}

func (gt *Gateway) Remove(id string) error {
    q := gt.decorate(bson.M{"_id": id})
    err := gt.db.Coll(gt.collectionName).Remove(q)
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

func handleError(db *s.MongoDB, err error) {
    if (err == io.EOF) {
        db.Session.Refresh()
    }
}