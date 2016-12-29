package resources

import (
    "gopkg.in/mgo.v2/bson"
    "github.com/supple/gorest/core"
    "github.com/supple/gorest/storage"
    "encoding/json"
    "io"
)

const REPO_DEVICE = "crm"

// --- ## Device model

type Device struct {
    CustomerBased `bson:",inline"`
    AppId      string        `json:"appId" bson:"appId" `
    AppToken   string        `json:"appToken" bson:"appToken"`
    AppVersion string         `json:"appVersion" bson:"appVersion"`
}

func (d *Device) IsValidForUpdate() {

}

// --- ## Device repository

type DeviceRP struct {
    gt *core.Gateway
    cc *core.CustomerContext
}

func NewDeviceRP(cc *core.CustomerContext) *DeviceRP {
    rp := &DeviceRP{cc: cc}
    db := storage.GetInstance(REPO_DEVICE)
    d := core.ContextDecorator(cc)
    gt := core.NewGateway(rp.CollectionName(), d, db)
    rp.gt = gt

    return rp
}

func (rp *DeviceRP) Create(model *Device) error {
    model.CustomerName = rp.cc.CustomerName
    model.AppId = rp.cc.AppId
    err := rp.ConstraintsValidation(model)
    if (err != nil) {
        return err
    }
    model.SetBasicFields()

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

func DeviceFromJson(data io.Reader) (*Device, error) {
    obj := &Device{}
    decoder := json.NewDecoder(data)
    if err := decoder.Decode(obj); err != nil {
        return nil, err
    }

    // validate

    return obj, nil
}

func (rp *DeviceRP) ConstraintsValidation(model *Device) (error) {
    var err error
    csRp := NewCustomerRP(rp.cc)
    _, err = csRp.FindOneByName(model.CustomerName)
    if err == core.ErrNotFound {
        // core.NewValidationError("customer", "Customer not found: ")
        return core.ErrorFrom(core.ErrNotFound,  "Customer not found")
    }
    if (err != nil) {
        return err
    }

    appRp := NewAppRP(rp.cc)
    _, err = appRp.FindOne(model.AppId)
    if err == core.ErrNotFound {
        // app_id, App id not set in api key
        return core.ErrorFrom(core.ErrNotFound,  "App id not set in api key")
    }
    if (err != nil) {
        return err
    }

    return err
}

func (rp DeviceRP) CollectionName() string {
    return "Device"
}
