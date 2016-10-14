package resources

import (
    "gopkg.in/mgo.v2/bson"
    s "github.com/supple/gorest/storage"
    lc "github.com/supple/gorest/utils"
    "reflect"
)

type Gateway struct {
    collectionName string
    cc *CustomerContext
}

func NewGateway(collectionName string, cc *CustomerContext) *Gateway {
    return &Gateway{collectionName: collectionName, cc: cc}
}

func (gt *Gateway) Insert(db *s.MongoDB, model interface{}) error {
    // generate id
    vDst := reflect.Indirect(reflect.ValueOf(model)).FieldByName("Id")
    if (vDst.Len() == 0) {
        vDst.SetString(lc.NewId())
    }

    err := db.Coll(gt.collectionName).Insert(model)
    return err
}

func (gt *Gateway) Remove(db *s.MongoDB, id string) (error) {
    q := bson.M{"_id": id, "customerName": gt.cc.CustomerName}
    err := db.Coll(gt.collectionName).Remove(q)
    return err
}

func (gt *Gateway) FindById(db *s.MongoDB, id string, result interface{}) error {
    q := bson.M{"_id": id, "customerName": gt.cc.CustomerName}
    err := db.Coll(gt.collectionName).Find(q).One(result)

    return err
}

func (gt *Gateway) FindOneBy(db *s.MongoDB, conditions bson.M, result interface{}) error {
    //conditions["customerName"] = gt.cc.CustomerName
    err := db.Coll(gt.collectionName).Find(conditions).One(result)
    return err
}