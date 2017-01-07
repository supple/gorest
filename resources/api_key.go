package resources

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "github.com/supple/gorest/core"
    "github.com/supple/gorest/model"
    "github.com/supple/gorest/storage"
    "fmt"
)

const API_KEY_FIELD string = "apiKey"
const API_KEY_REPO = "crm"

// ### -- ApiKey repo

type ApiKeyRP struct {
    gt *core.Gateway
    cc *core.CustomerContext
    db *storage.MongoDB
}

func NewApiKeyRP(cc *core.CustomerContext) *ApiKeyRP {
    rp := &ApiKeyRP{cc:cc}
    db := storage.GetInstance(API_KEY_REPO)
    d := core.EmptyDecorator()
    gt := core.NewGateway(rp.CollectionName(), d, db)

    rp.gt = gt

    return rp
}

func (rp *ApiKeyRP) Create(model *model.ApiKey) error {
    var err error
    model.CustomerName = rp.cc.CustomerName

     // validate
    customer, err := rp.ConstraintsValidation(model)
    if (err != nil) {
        return err
    }

    fmt.Println("OK CUST is")
    // create key if not set
    if (len(model.ApiKey) == 0) {
        model.ApiKey = fmt.Sprintf("%s-%s", core.RandString(24), customer.Hash)
    }
    err = rp.gt.Insert(model)
    fmt.Println(model)
    return err
}

func (rp *ApiKeyRP) FindOne(id string) (*model.ApiKey, error) {
    result := &model.ApiKey{}
    err := rp.gt.FindById(id, result)
    return result, err
}

func (rp *ApiKeyRP) FindOneBy(conditions bson.M) (*model.ApiKey, error) {
    result := &model.ApiKey{}
    err := rp.gt.FindOneBy(conditions, result)

    return result, err
}

func (rp *ApiKeyRP) Delete(id string) (error) {
    err := rp.gt.Remove(id)
    return err
}

func (rp *ApiKeyRP) ConstraintsValidation(model *model.ApiKey) (*model.Customer, error) {
    csRp := NewCustomerRP(rp.cc)
    c, err := csRp.FindOneByName(model.CustomerName)
    if (err == core.ErrNotFound) {
        return nil, ValidationError{Field: "customerName", Message: "Customer not found: "+model.CustomerName}
    }
    if (err != nil) {
        return nil, err
    }

    return c, err
}

func (rp ApiKeyRP) CollectionName() string {
    return "ApiKey"
}

func Auth(apiKey string, accessTo AccessTo) (*core.CustomerContext, error) {
    var cc *core.CustomerContext

    akRp := NewApiKeyRP(cc)
    ak, err := akRp.FindOneBy(bson.M{API_KEY_FIELD: apiKey})
    // @todo: hasAccess(accessTo)
    if err == core.ErrNotFound {
        return nil, core.ErrInvalidApiKey
    }

    if (err != nil) {
        return nil, err
    }

    if (ak != nil) {
        cc = &core.CustomerContext{}
        // copy to customer context
        cc.ApiKey = ak.ApiKey
        cc.CustomerName = ak.CustomerName
        cc.AppId = ak.AppId
    }

    return cc, err
}

func CreateApiKey(cc *core.CustomerContext) (*model.ApiKey, error) {
    akRp := NewApiKeyRP(cc)
    ak := &model.ApiKey{}
    ak.CustomerName = cc.CustomerName
    err := akRp.Create(ak)

    return ak, err
}

func (rp *ApiKeyRP) Install(db *storage.MongoDB) error {
    var index mgo.Index
    var err error

    // key in api key must be unique
    index = mgo.Index{
        Key: []string{"key"},
        Unique: true,
        DropDups: false,
        Background: true, // See notes.
        Sparse: true,
    }
    err = db.Coll(rp.CollectionName()).EnsureIndex(index)

    return err
}

