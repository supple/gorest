package resources

import (
    "gopkg.in/mgo.v2/bson"
    s "github.com/supple/gorest/storage"
)

type Device struct {
    CustomerBased `bson:",inline"`
    AppId      string        `json:"appId" bson:"appId" `
    AppToken   string        `json:"appToken" bson:"appToken"`
    AppVersion string         `json:"appVersion" bson:"appVersion"`
}

// --- ## Device repository

type DeviceRP struct {
    gt *Gateway
}

func NewDeviceRP() *DeviceRP {
    rp := &DeviceRP{}
    gt := &Gateway{collectionName: rp.CollectionName()}
    rp.gt = gt

    return rp
}

func (rp *DeviceRP) Create(db *s.MongoDB, model *Device) error {
    err := rp.ConstraintsValidation(db, model)
    if (err != nil) {
        return err
    }

    return rp.gt.Insert(db, model)
}

func (rp *DeviceRP) Update(db *s.MongoDB, id string, model *map[string]interface{}) error {
    err := db.Coll(rp.CollectionName()).Update(bson.M{"_id": id}, model)
    return err
}

func (rp *DeviceRP) FindOne(db *s.MongoDB, id string) (*Device, error) {
    result := &Device{}
    err := rp.gt.FindById(db, id, result)

    return result, err
}

func (rp *DeviceRP) FindOneBy(db *s.MongoDB, conditions bson.M) (*Device, error) {
    result := &Device{}
    err := rp.gt.FindOneBy(db, conditions, result)
    return result, err
}

func (rp *DeviceRP) Delete(db *s.MongoDB, id string) (error) {
    err := rp.gt.Remove(db, id)
    return err
}


func (rp *DeviceRP) ConstraintsValidation(db *s.MongoDB, model *Device) (error) {
    var err error
    csRp := NewCustomerRP()
    _, err = csRp.FindOneByName(db, model.CustomerName)
    if (err != nil) {
        return &ErrObjectNotFound{"Customer", model.CustomerName}
    }

    appRp := NewAppRP()
    _, err = appRp.FindOne(db, model.AppId)
    if (err != nil) {
        return &ErrObjectNotFound{"App", model.AppId}
    }

    return err
}

func (rp DeviceRP) CollectionName() string {
    return "Device"
}
