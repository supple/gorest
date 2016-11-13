package resources

import (
    "gopkg.in/mgo.v2/bson"
    "github.com/supple/gorest/core"
    s "github.com/supple/gorest/storage"
)

const REPO_DEVICE = "crm"

type Device struct {
    CustomerBased `bson:",inline"`
    AppId      string        `json:"appId" bson:"appId" `
    AppToken   string        `json:"appToken" bson:"appToken"`
    AppVersion string         `json:"appVersion" bson:"appVersion"`
}

// --- ## Device repository

type DeviceRP struct {
    gt *core.Gateway
    cc *core.CustomerContext
}

func NewDeviceRP(cc *core.CustomerContext) *DeviceRP {
    rp := &DeviceRP{cc: cc}
    db := s.GetInstance(REPO_DEVICE)
    gt := core.NewGateway(rp.CollectionName(), cc, db)
    rp.gt = gt

    return rp
}

func (rp *DeviceRP) Create(model *Device) error {
    err := rp.ConstraintsValidation(model)
    if (err != nil) {
        return err
    }

    return rp.gt.Insert(model)
}

func (rp *DeviceRP) Update(id string, model *map[string]interface{}) error {
    err := rp.gt.Update(id, model)
    return err
}

func (rp *DeviceRP) FindOne(id string) (*Device, error) {
    result := &Device{}
    err := rp.gt.FindById(id, result)

    return result, err
}

func (rp *DeviceRP) FindOneBy(conditions bson.M) (*Device, error) {
    result := &Device{}
    err := rp.gt.FindOneBy(conditions, result)
    return result, err
}

func (rp *DeviceRP) Delete(id string) (error) {
    err := rp.gt.Remove(id)
    return err
}


func (rp *DeviceRP) ConstraintsValidation(model *Device) (error) {
    var err error
    csRp := NewCustomerRP(rp.cc)
    _, err = csRp.FindOneByName(model.CustomerName)
    if (err != nil) {
        return &ErrObjectNotFound{"Customer", model.CustomerName}
    }

    appRp := NewAppRP(rp.cc)
    _, err = appRp.FindOne(model.AppId)
    if (err != nil) {
        return &ErrObjectNotFound{"App", model.AppId}
    }

    return err
}

func (rp DeviceRP) CollectionName() string {
    return "Device"
}
