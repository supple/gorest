package resources


import (
    "gopkg.in/mgo.v2/bson"
    s "github.com/supple/gorest/storage"
    lc "github.com/supple/gorest/utils"
    "reflect"
)

type Gateway struct {
    collectionName string
}

func (gt *Gateway) Insert(db *s.Mongo, model interface{}) error {

    // vDst := reflect.ValueOf(model).Elem().FieldByName("Id")
    vDst := reflect.Indirect(reflect.ValueOf(model)).FieldByName("Id")
    if (vDst.Len() == 0) {
        vDst.SetString(lc.NewId())
    }

    //fmt.Println(model)
    //if (len(model.(*Base).Id) == 0) {
    //    model.(*Base).Id = lc.NewId()
    //}
    err := db.Coll(gt.collectionName).Insert(model)
    return err
}

func (gt *Gateway) Remove(db *s.Mongo, id string) (error) {
    q := bson.M{"_id": id}
    err := db.Coll(gt.collectionName).Remove(q)
    return err
}

func (gt *Gateway) FindById(db *s.Mongo, id string, result interface{}) error {
    q := bson.M{"_id": id}
    err := db.Coll(gt.collectionName).Find(q).One(result)

    return err
}

func (gt *Gateway) FindOneBy(db *s.Mongo, conditions bson.M, result interface{}) error {
    err := db.Coll(gt.collectionName).Find(conditions).One(result)
    return err
}