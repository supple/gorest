package core

import (
    "gopkg.in/mgo.v2/bson"
    s "github.com/supple/gorest/storage"
    lc "github.com/supple/gorest/utils"
    "reflect"
    "gopkg.in/mgo.v2"
    "io"
    "time"
)

type Gateway struct {
    collectionName string
    cc *CustomerContext
    coll *mgo.Collection
    db *s.MongoDB
}

func NewGateway(collectionName string, cc *CustomerContext, db *s.MongoDB) *Gateway {
    return &Gateway{collectionName: collectionName, cc: cc, db: db}
}

func (gt *Gateway) getColl() {

}

func (gt *Gateway) Insert(model interface{}) error {
    // generate id
    vDst := reflect.Indirect(reflect.ValueOf(model)).FieldByName("Id")
    if (vDst.Len() == 0) {
        vDst.SetString(lc.NewId())
    }

    err := gt.db.Coll(gt.collectionName).Insert(model)
    handleError(gt.db, err)

    return err
}

func (gt *Gateway) Update(id string, model *map[string]interface{}) error {
    err := gt.db.Coll(gt.collectionName).Update(bson.M{"_id": id}, model)
    handleError(gt.db, err)

    return err
}

func (gt *Gateway) Remove(id string) (error) {
    q := bson.M{"_id": id, "customerName": gt.cc.CustomerName}
    err := gt.db.Coll(gt.collectionName).Remove(q)
    handleError(gt.db, err)

    return err
}

func (gt *Gateway) FindById(id string, result interface{}) error {
    q := bson.M{"_id": id, "customerName": gt.cc.CustomerName}
    err := gt.db.Coll(gt.collectionName).Find(q).One(result)
    handleError(gt.db, err)

    return err
}

func (gt *Gateway) FindOneBy(conditions bson.M, result interface{}) error {
    conditions["customerName"] = gt.cc.CustomerName
    err := gt.db.Coll(gt.collectionName).Find(conditions).One(result)
    handleError(gt.db, err)

    return err
}

// Find element without customer context
func (gt *Gateway) FindInsecureOneBy(conditions bson.M, result interface{}) error {
    time.Sleep(time.Second)
    err := gt.db.Coll(gt.collectionName).Find(conditions).One(result)
    handleError(gt.db, err)
    return err
}

func handleError(db *s.MongoDB, err error) {
    if (err == io.EOF) {
        db.Session.Refresh()
    }
}