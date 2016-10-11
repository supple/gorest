package resources


import (
    "gopkg.in/mgo.v2/bson"
    s "github.com/supple/gorest/storage"
)

type Device struct {
    Base
    Id           string `json:"id" bson:"_id"`
    CustomerName string `json:"customerName" bson:"customerName" validate:"required"`
    AppId        string        `json:"name" bson:"name" `
    AppToken     string        `json:"appToken" bson:"appToken"`
    AppVersion   int32       `json:"appVersion" bson:"appVersion"`
}

// --- ## Device repository

type DeviceRP struct{
    gt *Gateway
}

func NewDeviceRP() *DeviceRP {
    rp := &DeviceRP{}
    gt := &Gateway{collectionName: rp.CollectionName()}
    rp.gt = gt

    return rp
}

func (rp *DeviceRP) Create(db *s.Mongo, model *Customer) error {
    return rp.gt.Insert(db, model)
}

func (rp *DeviceRP) Update(db *s.Mongo, id string, model *map[string]interface{}) error {
    err := db.Coll(rp.CollectionName()).Update(bson.M{"_id": id}, model)
    return err
}

func (rp *DeviceRP) FindOne(db *s.Mongo, id string) (*Device, error) {
    result := &Device{}
    err := rp.gt.FindById(db, id, result)

    return result, err
}

func (rp *DeviceRP) FindOneBy(db *s.Mongo, conditions bson.M) (*Device, error) {
    result := &Device{}
    err := rp.gt.FindOneBy(db, conditions, result)
    return result, err
}

func (rp *DeviceRP) Delete(db *s.Mongo, id string) (error) {
    err := rp.gt.Remove(db, id)
    return err
}

func (rp DeviceRP) CollectionName() string {
    return "Device"
}
